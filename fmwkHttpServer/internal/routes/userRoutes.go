package routes

import (
	"fmwkHttpServer/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	r fiber.Router,
	userHandler *handlers.UserHandler,
) {
	r.Get("/users", userHandler.GetUsers)

	r.Get("/users/:id", userHandler.GetOneUser)
}
