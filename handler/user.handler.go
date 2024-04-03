package handler

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/model/response"
	"go-fiber-gorm/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {

	var user request.UserCreateRequest

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	// Check email exist
	var userExist entity.User
	database.DB.Unscoped().Where("email =?", user.Email).First(&userExist)

	if userExist.ID != 0 {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "email already exist",
		})
	}

	if err := request.ValidateUserCreateRequest(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(), // Mengembalikan pesan kesalahan validasi

		})
	}

	// Setelah validasi, konversi ke struktur data User
	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	hashPassword, err := utils.HashString(user.Password)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	newUser.Password = hashPassword

	result := database.DB.Debug().Create(&newUser)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created",
		"data":    newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	userResponse := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Phone:     user.Phone,
		Address:   user.Address,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    userResponse,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User

	userId := ctx.Params("id")

	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// Update User Data
	if strings.TrimSpace(userRequest.Name) != "" {
		user.Name = userRequest.Name
	}
	user.Phone = userRequest.Phone
	user.Address = userRequest.Address

	result := database.DB.Debug().Model(&user).Updates(user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "user updated",
		"data":    user,
	})
}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	userRequest := new(request.UserEmailRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User

	userId := ctx.Params("id")

	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// Check email exist
	var userExist entity.User
	database.DB.Unscoped().Where("email =?", userRequest.Email).First(&userExist)

	if userExist.ID != 0 {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "email already exist",
		})
	}

	// Update User Data
	user.Email = userRequest.Email

	result := database.DB.Debug().Model(&user).Updates(user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user updated",
		"data":    user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User
	err := database.DB.First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	result := database.DB.Debug().Delete(&user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user deleted",
	})

}
