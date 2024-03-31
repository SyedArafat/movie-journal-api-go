package services

import (
	"golang.org/x/crypto/bcrypt"
	"movie-journal-api/internal/logger"
	"movie-journal-api/internal/model"
	"movie-journal-api/internal/repository/databaseRepo"
	"movie-journal-api/internal/validator"
)

func (as *ApplicationService) AddUser(user validator.User) error {
	repo := databaseRepo.NewApplicationRepo(as.Application)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		logger.Error("PASSWORD_HASHING_ERROR", "Failed to hash password", "", err)
	}

	u := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	return repo.AddUser(u)
}
