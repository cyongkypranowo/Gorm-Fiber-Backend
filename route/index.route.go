package route

import (
	"go-fiber-gorm/config"
	"go-fiber-gorm/handler"
	"go-fiber-gorm/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectPath+"/public/assets")

	r.Post("/login", handler.LoginHandler)

	r.Get("/user", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDelete)
}
