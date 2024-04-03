package handler

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(ctx *fiber.Ctx) error {
	var loginRequest request.LoginRequest

	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := request.ValidateLoginRequest(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			// Mengembalikan pesan kesalahan validasi
		})
	}

	var user entity.User
	err := database.DB.First(&user, "email=?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	// Check Validation Password
	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "wrong credentials",
		})
	}

	//Generate Token
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	claims["role"] = "admin"

	token, err := utils.GenerateToken(&claims)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login success",
		"token":   token,
	})

}
