package utils

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSingleFile(ctx *fiber.Ctx) error {
	// Handle File
	file, errFile := ctx.FormFile("file")
	if errFile != nil {
		log.Println("Error:", errFile.Error())
	}

	var filename string

	if file != nil {
		filename = file.Filename
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/assets/%s", filename))
		if errSaveFile != nil {
			log.Println("Fail to save file: ", errSaveFile)
		}
	} else {
		log.Println("Nothing to save")
	}

	ctx.Locals("filename", filename)
	return ctx.Next()
}

func HandleMultipleFiles(ctx *fiber.Ctx) error {
	// Handle File
	form, errFile := ctx.MultipartForm()
	if errFile != nil {
		log.Println("Error:", errFile.Error())
	}

	files := form.File["files"]
	var filenames []string

	for i, file := range files {
		var filename string
		if file != nil {
			filename = fmt.Sprintf("%d-%v", i, file.Filename) // Handle jika nama file sama, akan di tambahkan menggunakan index
			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/assets/%s", filename))
			if errSaveFile != nil {
				log.Println("Fail to save file into public/assets directory: ", errSaveFile)
			}
		} else {
			log.Println("Nothing file to uploading")
		}

		if filename != "" {
			filenames = append(filenames, filename)
		}
	}

	ctx.Locals("files", filenames)

	return ctx.Next()
}
