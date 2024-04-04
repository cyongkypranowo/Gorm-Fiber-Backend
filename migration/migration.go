package migration

import (
	"fmt"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"log"
)

func RunMigrations() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Category{}, &entity.Photo{})
	if err != nil {
		log.Println((err))
	}
	fmt.Println("migration successfully migrated")
}
