package route

import (
	"go-fiber-gorm/config"
	"go-fiber-gorm/handler"
	"go-fiber-gorm/middleware"
	"go-fiber-gorm/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Static("/public", config.ProjectPath+"/public/assets")

	r.Post("/login", handler.LoginHandler)

	r.Get("/user", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/user/:id", middleware.Auth, handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", middleware.Auth, handler.UserHandlerUpdate)
	r.Put("/user/:id/update-email", middleware.Auth, handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", middleware.Auth, handler.UserHandlerDelete)

	r.Post("/book", utils.HandleSingleFile("Book"), handler.BookHandlerCreate)

	r.Post("/photo", utils.HandleMultipleFile("Photo"), handler.PhotoHandlerCreate)

}
