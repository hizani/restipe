package recipeservice

import (
	"errors"
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

func (r *RecipeService) GetAll(recipe model.GetAllRecipesReq) ([]model.AllRecipeResp, error) {
	return r.storage.GetAll(recipe)
}

func (r *RecipeService) GetById(recipeId int) (model.RecipeResp, error) {
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

func (r *RecipeService) RemoveStepFromRecipe(userId, recipeId, number int) error {
	return r.storage.RemoveStepFromRecipe(userId, recipeId, number)
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

func (r *RecipeService) RateRecipe(userId, recipeId int, rating model.RateReq) (int, error) {
	if rating.Rating < 1 || rating.Rating > 5 {
		return 0, errors.New("rating should be >=1 and <=5")
	}
	return r.storage.RateRecipe(userId, recipeId, rating)
}

func (r *RecipeService) RerateRecipe(userId, recipeId int, rating model.RateReq) error {
	if rating.Rating < 1 || rating.Rating > 5 {
		return errors.New("rating should be >=1 and <=5")
	}
	return r.storage.RerateRecipe(userId, recipeId, rating)
}

func (r *RecipeService) GetRecipeImgFilename(recipeId int) (*string, error) {
	filename, err := r.storage.GetRecipeImgFilename(recipeId)
	if err != nil {
		return filename, err
	}
	if filename == nil || *filename == "" {
		return filename, errors.New("recipe has no image")
	}
	return filename, err
}

func (r *RecipeService) UpdateRecipeImgFilename(userId, recipeId int, filename *string) (*string, error) {
	return r.storage.UpdateRecipeImgFilename(userId, recipeId, filename)

}

func (r *RecipeService) GetStepImgFilename(recipeId, number int) (*string, error) {
	filename, err := r.storage.GetStepImgFilename(recipeId, number)
	if err != nil {
		return filename, err
	}
	if filename == nil || *filename == "" {
		return filename, errors.New("step has no image")
	}
	return filename, err
}

func (r *RecipeService) UpdateStepImgFilename(userId, recipeId, number int, filename *string) (*string, error) {
	return r.storage.UpdateStepImgFilename(userId, recipeId, number, filename)
}
