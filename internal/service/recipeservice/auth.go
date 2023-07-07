package recipeservice

import (
	"crypto/sha256"
	"restipe/internal/model"
	"restipe/internal/storage"
)

type AuthService struct {
	storage storage.Authorization
}

func NewAuthService(storage storage.Authorization) *AuthService {
	return &AuthService{storage}
}

func (s *AuthService) SignupUser(user model.SignupUser) (int, error) {
	user.Password = hashPassword(user.Password)
	return s.storage.SignupUser(user)
}

func (s *AuthService) SigninUser(user model.SigninUser) (int, error) {
	user.Password = hashPassword(user.Password)
	return s.storage.SigninUser(user)
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return string(hash.Sum(nil))

}
