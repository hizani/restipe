package storage

import (
	"restipe/internal/model"
	"restipe/internal/storage/sqldb"
	"restipe/internal/storage/sqldb/postgres"
)

type Authorization interface {
	SignupUser(user model.SignupUserReq) (int, error)
	SigninUser(user model.SigninUserReq) (int, error)

	Close() error
}

type Recipe interface {
	Create(userId int, recipe model.CreateRecipeReq) (int, error)
	Delete(userId, recipeId int) error
	Update(userId, recipeId int, recipe model.UpdateRecipeReq) error
	GetAll(recipe model.GetAllRecipesReq) ([]model.AllRecipeResp, error)
	GetById(recipeId int) (model.RecipeResp, error)
	GetAllIngredientsFromRecipe(recipeId int) ([]model.Ingredient, error)
	GetAllStepsFromRecipe(recipeId int) ([]model.Step, error)
	AddStepToRecipe(userId, recipeId int, step model.AddStepReq) (int, error)
	AddIngredientToRecipe(userId, recipeId int, ingredient model.AddIngredientReq) (int, error)
	RemoveStepFromRecipe(userId, recipeId, stepId int) error
	RemoveIngredientFromRecipe(userId, recipeId, ingredientId int) error

	Close() error
}

type Storage struct {
	Authorization
	Recipe
}

func NewPostgres(cfg sqldb.Config) (*Storage, error) {
	db, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}
	return &Storage{
		Authorization: postgres.NewAuthStoarge(db),
		Recipe:        postgres.NewRecipeStorage(db),
	}, err
}

func (s *Storage) Close() error {
	if err := s.Authorization.Close(); err != nil {
		return err
	}
	return s.Recipe.Close()
}
