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
	Delete(userId, recipeId int) error
	Update(userId, recipeId int, recipe model.UpdateRecipeReq) error
	AddStepToRecipe(userId, recipeId int, step model.AddStepReq) (int, error)
	AddIngredientToRecipe(userId int, recipeId int, ingredient model.AddIngredientReq) (int, error)
	GetAll(recipe model.GetAllRecipesReq) ([]model.AllRecipeResp, error)
	GetById(recipeId int) (model.RecipeResp, error)
	GetAllIngredientsFromRecipe(recipeId int) ([]model.Ingredient, error)
	GetAllStepsFromRecipe(recipeId int) ([]model.Step, error)
	RemoveStepFromRecipe(userId, recipeId, number int) error
	RemoveIngredientFromRecipe(userId, recipeId, ingredientId int) error
	RateRecipe(userId, recipeId int, rating model.RateReq) (int, error)
	RerateRecipe(userId, recipeId int, rating model.RateReq) error
	GetRecipeImgFilename(recipeId int) (*string, error)
	UpdateRecipeImgFilename(userId, recipeId int, filename *string) (*string, error)
	GetStepImgFilename(recipeId, number int) (*string, error)
	UpdateStepImgFilename(userId, recipeId, number int, filename *string) (*string, error)
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
