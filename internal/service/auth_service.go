package service

import (
	"errors"
	"fin_tracker/internal/models"
	"fin_tracker/internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret-key") // для dev-режима; вынесем позже в конфиг

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo}
}

func (s *AuthService) Register(email, password, fullName string) (string, error) {
	_, err := s.userRepo.FindByEmail(email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := models.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: string(hash),
		FullName:     fullName,
	}

	if err := s.userRepo.Create(&user); err != nil {
		return "", err
	}

	return generateJWT(user.ID)
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return generateJWT(user.ID)
}

func generateJWT(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
