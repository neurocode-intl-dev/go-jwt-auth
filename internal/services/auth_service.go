package services

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/rahulreghupdm/tashub-project/internal/repository"
	"github.com/rahulreghupdm/tashub-project/internal/utils"
)

func Register(email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	// Use repository instead of direct DB calls
	return repository.CreateUser(email, string(hashed))
}

func Login(email, password string) (string, string, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}

	if user == nil {
		return "", "", err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", "", err
	}

	access, _ := utils.GenerateAccessToken(user.ID)
	refresh, _ := utils.GenerateRefreshToken(user.ID)

	return access, refresh, nil
}
