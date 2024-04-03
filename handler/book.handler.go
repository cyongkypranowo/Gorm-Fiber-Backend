package handler

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"log"

	"github.com/gofiber/fiber/v2"
)

func BookHandlerCreate(ctx *fiber.Ctx) error {

	var book request.BookCreateRequest

	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	// Validasi Request
	if err := request.ValidateBookCreateRequest(&book); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(), // Mengembalikan pesan kesalahan validasi

		})
	}

	// Handle File
	ctx.FormFile("cover")

	file, errFile := ctx.FormFile("cover")
	if errFile != nil {
		log.Println("Error:", errFile.Error())
	}

	var filename string

	if file != nil {
		filename = file.Filename
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/assets/%s", filename))
		if errSaveFile != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": errSaveFile.Error(),
			})
		}
	} else {
		log.Println("Nothing to save")
	}

	// Setelah validasi, konversi ke struktur data User
	newBook := entity.Book{
		Title:  book.Title,
		Cover:  filename,
		Author: book.Author,
	}

	result := database.DB.Debug().Create(&newBook)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": result.Error,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created",
		"data":    newBook,
	})
}
