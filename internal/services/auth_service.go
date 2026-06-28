package services

import (
	"foodiego/internal/config"
	"foodiego/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db *config.DB
}

func (s *AuthService) RegisterUser(user *models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	user = &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	s.db.Create(&user)
	return nil
}
