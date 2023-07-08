package service

import (
	"restipe/internal/model"
	"restipe/internal/service/recipeservice"
	"restipe/internal/storage"
)

type Authorization interface {
	SignupUser(user model.SignupUserReq) (int, error)
	SigninUser(user model.SigninUserReq) (string, error)
	Authorize(token string) (int, error)
}

type Recipe interface {
	Create(userId int, recipe model.CreateRecipeReq) (int, error)
	GetAll(recipe model.GetAllRecipesReq) ([]model.Recipe, error)
	GetById(recipeId int) (model.Recipe, error)
	GetAllIngredientsFromRecipe(recipeId int) ([]model.Ingredient, error)
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
