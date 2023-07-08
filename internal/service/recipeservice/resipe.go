package recipeservice

import (
	"restipe/internal/model"
	"restipe/internal/storage"
)

type RecipeService struct {
	storage storage.Recipe
}

func NewRecipeService(storage storage.Recipe) *RecipeService {
	return &RecipeService{storage}
}

func (r *RecipeService) Create(userId int, recipe model.CreateRecipeReq) (int, error) {
	return r.storage.Create(userId, recipe)
}

func (r *RecipeService) GetAll(recipe model.GetAllRecipesReq) ([]model.Recipe, error) {
	return r.storage.GetAll(recipe)
}

func (r *RecipeService) GetById(recipeId int) (model.Recipe, error) {
	return r.storage.GetById(recipeId)
}

func (r *RecipeService) GetAllIngredientsFromRecipe(recipeId int) ([]model.Ingredient, error) {
	return r.storage.GetAllIngredientsFromRecipe(recipeId)
}

func (r *RecipeService) GetAllStepsFromRecipe(recipeId int) ([]model.Step, error) {
	return r.storage.GetAllStepsFromRecipe(recipeId)
}

func (r *RecipeService) AddStepToRecipe(userId, recipeId int, step model.AddStepReq) (int, error) {
	return r.storage.AddStepToRecipe(userId, recipeId, step)

}

func (r *RecipeService) RemoveStepFromRecipe(userId, recipeId, stepId int) error {
	return r.storage.RemoveStepFromRecipe(userId, recipeId, stepId)
}
