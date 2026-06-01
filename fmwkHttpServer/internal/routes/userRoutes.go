package routes

import (
	"fmwkHttpServer/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	userHandler *handlers.UserHandler,
) {
	app.Get("/users", userHandler.GetUsers)

	app.Get("/users/:id", userHandler.GetOneUser)
}
