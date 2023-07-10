package model

type AllRecipeResp struct {
	Id          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Description string  `json:"description" db:"description"`
	Author      int     `json:"author" db:"author"`
	Duration    int64   `json:"duration" db:"duration"`
	AvgRating   float32 `json:"avg_rating" db:"avg_rating"`
}

type RecipeResp struct {
	Id          int          `json:"id" db:"id"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	Author      int          `json:"author" db:"author"`
	Duration    int64        `json:"duration" db:"duration"`
	AvgRating   float32      `json:"avg_rating" db:"avg_rating"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []Step       `json:"steps"`
}

type CreateRecipeReq struct {
	Name        string             `json:"name" binding:"required"`
	Description string             `json:"description"`
	Ingredients []AddIngredientReq `json:"ingredients"`
	Steps       []CreateStepReq    `json:"steps"`
}

type UpdateRecipeReq struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type GetAllRecipesReq struct {
	IngredientFilter []int    `json:"ingredient_filter" example:"1"`
	DurationFilter   *int64   `json:"duration_filter" example:"7200"`
	RatingFilter     *float32 `json:"rating_filter" example:"4.5"`
	DurationSort     *string  `json:"duration_sort" example:"DESC"`
	RatingSort       *string  `json:"rating_sort" example:"ASC"`
	Author           *int     `json:"author" example:"1"`
}
