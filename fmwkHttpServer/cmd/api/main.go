package main

import (
	"fmwkHttpServer/internal/database"
	"fmwkHttpServer/internal/handlers"
	"fmwkHttpServer/internal/repository"
	"fmwkHttpServer/internal/routes"
	"fmwkHttpServer/internal/services"

	"github.com/gofiber/fiber/v2"

	"log"
)

func main() {
	db := database.ConnectDB()

	repo := repository.NewUserRepository(db)

	userService := service.NewUserService(repo)

	handler := handlers.NewUserHandler(userService)

	app := fiber.New()

	routes.SetupRoutes(app, handler)

	log.Println("server up at 8080")

	app.Listen(":8080")

}
