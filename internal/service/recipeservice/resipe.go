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

func (r *RecipeService) Delete(userId, recipeId int) error {
	return r.storage.Delete(userId, recipeId)

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

func (r *RecipeService) RemoveIngredientFromRecipe(userId, recipeId, ingredientId int) error {
	return r.storage.RemoveIngredientFromRecipe(userId, recipeId, ingredientId)
}

func (r *RecipeService) AddIngredientToRecipe(userId int, recipeId int, ingredient model.AddIngredientReq) (int, error) {
	return r.storage.AddIngredientToRecipe(userId, recipeId, ingredient)
}

func (r *RecipeService) Update(userId, recipeId int, recipe model.UpdateRecipeReq) error {
	return r.storage.Update(userId, recipeId, recipe)

}
