package dbModels

import (
	"strings"

	validationLibs "github.com/wildanpurnomo/go-persistent-chat/server/libs/validation"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID    int64  `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) ValidateRegistrationInputs() bool {
	return validationLibs.IsValidUsername(u.Username) && validationLibs.IsValidPassword(u.Password)
}

func (u *User) TrimUsername() {
	u.Username = strings.TrimSpace(u.Username)
}

func (u *User) HashPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)

	return nil
}

func (u *User) VerifyPassword(passwordInput string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwordInput)); err != nil {
		return false
	}

	return true
}
