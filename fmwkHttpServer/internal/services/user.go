package service

import (
	"fmwkHttpServer/internal/models"
	"fmwkHttpServer/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(
	repo *repository.UserRepository,
) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.repo.GetUsers()
}

func (s *UserService) GetOneUser(id int) (*model.User, error) {
	return s.repo.GetOneUser(id)
}

func (s *UserService) AddUser(user model.User) (*model.User, error) {
	return s.repo.AddUser(user)
}
