package service

import (
	"fin_tracker/internal/model"
	"fin_tracker/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) GetByID(id string) (*model.User, error) {
	return s.userRepo.FindByID(id)
}
