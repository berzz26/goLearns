package handlers

import (
	"fmwkHttpServer/internal/models"

	"fmwkHttpServer/internal/services"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(

	service *service.UserService,

) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUsers() ([]model.User, error) {
	return h.service.GetUsers()
}
func (h *UserHandler) GetOneUser(id int) (*model.User, error) {
	return h.service.GetOneUser(id)
}
