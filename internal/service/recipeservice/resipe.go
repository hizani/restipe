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

func (r *RecipeService) Create(userId int, recipe model.CreateRecipe) (int, error) {
	return r.storage.Create(userId, recipe)
}
