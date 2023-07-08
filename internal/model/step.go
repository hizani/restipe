package model

type CreateStepReq struct {
	Id          int    `json:"id"`
	RecipeId    int    `json:"recipe_id" db:"recipe_id"`
	Number      int    `json:"number" db:"number"`
	Description string `json:"description" binding:"required" db:"description"`
	Duration    int64  `json:"duration" binding:"required" db:"duration"`
}

type Step struct {
	Id          int    `json:"id" db:"id"`
	Number      int    `json:"number" db:"number"`
	Description string `json:"descriprion" db:"description"`
	Duration    int64  `json:"duration" db:"duration"`
}
