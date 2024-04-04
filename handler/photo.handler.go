package handler

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"

	"github.com/gofiber/fiber/v2"
)

func PhotoHandlerCreate(ctx *fiber.Ctx) error {
	var photo request.PhotoCreateRequest

	if err := ctx.BodyParser(&photo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	// Validasi Request
	if err := request.ValidatePhotoCreateRequest(&photo); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(), // Mengembalikan pesan kesalahan validasi
		})
	}

	filesInterface := ctx.Locals("files")
	filenames, ok := filesInterface.([]string)
	if !ok {
		// Handle kesalahan jika konversi gagal
		fmt.Println("Tidak dapat mengonversi nilai 'files' ke slice string")
		return nil
	}

	// Check jika multiple files requires
	if len(filenames) == 0 {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "file upload required",
		})
	}

	processed := make(map[string]bool)
	var photos []entity.Photo

	// Iterasi melalui setiap nilai dalam slice filenames
	for _, v := range filenames {
		// Mengecek apakah nilai sudah diproses sebelumnya
		if _, ok := processed[v]; !ok {
			// Jika nilai belum diproses sebelumnya, tambahkan ke dalam map dan slice
			processed[v] = true
			photos = append(photos, entity.Photo{
				Name:       v,
				CategoryId: 1,
			})
		}
	}

	// Lakukan operasi bulk insert ke dalam database
	if err := database.DB.Debug().Create(&photos).Error; err != nil {
		// Handle error jika terjadi
		fmt.Println("Failed to bulk insert photos:", err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "photo created",
		"data":    "berhasil",
	})
}
