package handler

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {
	var loginRequest request.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err,
			"data":    nil,
		})
	}

	if err := request.ValidateLoginRequest(&loginRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil, // Mengembalikan pesan kesalahan validasi
		})
	}

	var user entity.User
	err := database.DB.First(&user, "email=?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "wrong credentials",
			"data":    nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "login success",
		"data": fiber.Map{
			"token": "secret",
			"user":  user,
		},
	})

}
