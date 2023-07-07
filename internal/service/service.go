package service

import (
	"restipe/internal/model"
	"restipe/internal/service/recipeservice"
	"restipe/internal/storage"
)

type Authorization interface {
	SignupUser(user model.SignupUser) (int, error)
	SigninUser(user model.SigninUser) (int, error)
}

type Recipe interface {
}

type Service struct {
	Authorization
	Recipe
}

func New(storage *storage.Storage) *Service {
	return &Service{
		Authorization: recipeservice.NewAuthService(storage),
	}
}
