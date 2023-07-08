package model

type Recipe struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Author      int    `json:"author" db:"author"`
}

type CreateRecipeReq struct {
	Name        string          `json:"name" binding:"required"`
	Description string          `json:"description"`
	Ingredients []Ingredient    `json:"ingredients"`
	Steps       []CreateStepReq `json:"steps"`
}

type UpdateRecipeReq struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type GetAllRecipesReq struct {
	IngredientFilter []int  `json:"ingredient_filter"`
	DurationSort     string `json:"duration_sort"`
	Author           int    `json:"author"`
}
