package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const defaultPathAssetImage = "./public/uploads"

func HandleSingleFile(module_name string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Handle File
		file, errFile := ctx.FormFile("file")

		if errFile != nil {
			log.Println("Error:", errFile.Error())
		}

		var filename string

		if file != nil {
			filename = file.Filename
			// Get Path file
			path, ext := getPathAsset(filename)
			if path == "" {
				log.Println("file not found")
				return nil
			}

			// generate filename
			fGenerator := uuid.New().String()

			path = defaultPathAssetImage + "/" + path + "/" + module_name + "_" + fGenerator + ext
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
}

func HandleMultipleFile(module_name string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Handle File
		form, errFile := ctx.MultipartForm()
		if errFile != nil {
			log.Println("Error:", errFile.Error())
			return errFile
		}

		files := form.File["files"]
		var filenames []string

		for _, file := range files {
			var filename string
			if file != nil {
				filename = file.Filename
				// Get Path file
				path, ext := getPathAsset(filename)
				if path == "" {
					log.Println("file not found")
					return nil
				}

				// generate filename
				fGenerator := uuid.New().String()

				path = defaultPathAssetImage + "/" + path + "/" + module_name + "_" + fGenerator + ext
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
}

func HandleRemoveFile(filename string) error {

	path, _ := getPathAsset(filename)
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

func getPathAsset(filename string) (string, string) {
	fmt.Println(filename)
	if filename == "" {
		return "", ""
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

	return path, ext
}
