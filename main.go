package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/migration"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize database
	database.DatabaseInit()
	migration.RunMigrations()

	app := fiber.New()

	// Initil route
	route.RouteInit(app)

	app.Listen(":8080")
}
