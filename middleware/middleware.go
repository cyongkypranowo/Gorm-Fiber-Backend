package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Auth(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	// Periksa apakah header Authorization kosong
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "missing authorization header",
		})
	}

	// Periksa apakah header Authorization memiliki format yang benar
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid authorization header format",
		})
	}

	// Dapatkan nilai token dari bagian kedua (indeks 1)
	token := parts[1]

	// Periksa apakah token sesuai dengan nilai yang diharapkan
	if token != "secret" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	// Lanjutkan ke handler berikutnya jika token valid
	return ctx.Next()
}

func PermisionCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}
