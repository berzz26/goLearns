package main

import (
	"fmwkHttpServer/internal/database"
	"fmwkHttpServer/internal/handlers"
	"log"
)

func main() {
	db := database.ConnectDB()

	users,err := handlers.GetUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(users)
}
