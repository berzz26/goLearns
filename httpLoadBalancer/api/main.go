package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	api.Get("/", LogRequests)
	
	app.Listen(":3000")

}

func LogRequests(c *fiber.Ctx) error {
	log.Println("Request recieved at ", time.Now())
	return nil
}
