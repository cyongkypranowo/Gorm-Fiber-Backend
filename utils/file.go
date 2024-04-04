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
