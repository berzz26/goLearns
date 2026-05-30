package handlers

import (
	"database/sql"
	"fmwkHttpServer/internal/models"
	"fmwkHttpServer/internal/repository"
	"fmwkHttpServer/internal/services"
)

func GetUsers(db *sql.DB) ([]model.User, error) {

	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)

	users,err := service.GetUsers()

	return users,err


}
