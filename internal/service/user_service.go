package service

import (
	"fin_tracker/internal/models"
	"fin_tracker/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) GetByID(id string) (*models.User, error) {
	return s.userRepo.FindByID(id)
}
