package model

type Recipe struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Author      int    `json:"author" db:"author"`
}

type GetRecipeReq struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Author      int    `json:"author" db:"author"`
}

type CreateRecipeReq struct {
	Name        string          `json:"name" binding:"required"`
	Description string          `json:"description"`
	Ingredients []Ingredient    `json:"ingredients" binding:"required"`
	Steps       []CreateStepReq `json:"steps" binding:"required"`
}

type GetAllRecipesReq struct {
	IngredientFilter []int  `json:"ingredient_filter"`
	DurationSort     string `json:"duration_sort"`
	Author           int    `json:"author"`
}
