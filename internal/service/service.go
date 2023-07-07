package service

import (
	"restipe/internal/model"
	"restipe/internal/service/recipeservice"
	"restipe/internal/storage"
)

type Authorization interface {
	SignupUser(user model.SignupUser) (int, error)
	SigninUser(user model.SigninUser) (string, error)
	Authorize(token string) (int, error)
}

type Recipe interface {
	Create(userId int, recipe model.CreateRecipe) (int, error)
}

type Service struct {
	Authorization
	Recipe
}

func New(storage *storage.Storage, jwttoken string) *Service {
	return &Service{
		Authorization: recipeservice.NewAuthService(storage.Authorization, jwttoken),
		Recipe:        recipeservice.NewRecipeService(storage.Recipe),
	}
}
