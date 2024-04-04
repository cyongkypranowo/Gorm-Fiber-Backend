package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

const defaultPathAssetImage = "./public/uploads"

func HandleSingleFile(ctx *fiber.Ctx) error {
	// Handle File
	file, errFile := ctx.FormFile("file")
	if errFile != nil {
		log.Println("Error:", errFile.Error())
	}

	var filename string

	// Get Path file
	path := getPathAsset(filename)
	if path == "" {
		log.Println("file not found")
		return nil
	}
	path = defaultPathAssetImage + "/" + path + "/" + filename

	if file != nil {
		filename = file.Filename
		errSaveFile := ctx.SaveFile(file, path)
		if errSaveFile != nil {
			log.Println("Fail to save file: ", errSaveFile)
		}
	} else {
		log.Println("Nothing to save")
	}

	ctx.Locals("filename", filename)
	return ctx.Next()
}

func HandleMultipleFile(ctx *fiber.Ctx) error {
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
			// Get Path file
			path := getPathAsset(filename)
			if path == "" {
				log.Println("file not found")
				return nil
			}
			path = defaultPathAssetImage + "/" + path + "/" + filename
			errSaveFile := ctx.SaveFile(file, path)
			if errSaveFile != nil {
				log.Println("Fail to save file into public/uploads directory: ", errSaveFile)
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

func HandleRemoveFile(filename string) error {

	path := getPathAsset(filename)
	if path == "" {
		log.Println("file not found")
		return nil
	}

	path = defaultPathAssetImage + "/" + path + "/" + filename

	err := os.Remove(path)
	if err != nil {
		log.Printf("Failed to remove file : %v", err)
		return err
	}

	return nil
}

func getPathAsset(filename string) string {
	if filename == "" {
		return ""
	}

	// get ext of filename
	ext := filepath.Ext(filename)
	var path string
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".ico", ".jfif", ".svg", ".webp":
		path = "images"
	case ".pdf", ".doc", ".docx", ".txt", ".odt", ".rtf", ".xls", ".xlsx":
		path = "documents"
	case ".mp4", ".avi", ".mov", ".mkv", ".wmv", ".webm":
		path = "videos"
	case ".mp3", ".wav", ".ogg":
		path = "musics"
	default:
		path = "miscellaneous" // Jika ekstensi tidak dikenali
	}

	return path
}
