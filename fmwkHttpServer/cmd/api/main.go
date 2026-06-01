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
	//production grade router setup
	// "/api/v1.."
	app := fiber.New()
	api := app.Group("/api")
	// seperate v1 group so that versioning could be easily managed
	v1 := api.Group("/v1")

	routes.SetupRoutes(v1, handler)

	log.Println("server up at 8080")

	app.Listen(":8080")

}
