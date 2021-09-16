package controllers

import (
	"errors"
	"strings"

	"github.com/wildanpurnomo/go-persistent-chat/server/db"
	dbModels "github.com/wildanpurnomo/go-persistent-chat/server/db/models"
)

func Authenticate(userModel *dbModels.User) error {
	if err := db.AppRepository.GetUserById(userModel); err != nil {
		return err
	}

	return nil
}

func Register(userModel *dbModels.User) error {
	userModel.TrimUsername()

	if isValid := userModel.ValidateRegistrationInputs(); !isValid {
		return errors.New("Invalid username or password")
	}

	if err := userModel.HashPassword(); err != nil {
		return errors.New("An unexpected problem occurs")
	}

	if err := db.AppRepository.CreateNewUser(userModel); err != nil {
		return err
	}

	return nil
}

func Login(username string, password string) (*dbModels.User, error) {
	var userModel dbModels.User

	userModel.Username = strings.TrimSpace(username)

	if err := db.AppRepository.SearchUserByUsername(&userModel); err != nil {
		return nil, err
	}

	if isValidPassword := userModel.VerifyPassword(password); !isValidPassword {
		return nil, errors.New("Invalid credentials or user not found")
	}

	return &userModel, nil
}
