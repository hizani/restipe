package model

type CreateStepReq struct {
	RecipeId    int    `json:"recipe_id" db:"recipe_id"`
	Description string `json:"description" binding:"required" db:"description"`
	Duration    int64  `json:"duration" binding:"required" db:"duration"`
}

type Step struct {
	Id          int    `json:"id" db:"id"`
	Number      int    `json:"number" db:"number"`
	Description string `json:"descriprion" db:"description"`
	Duration    int64  `json:"duration" db:"duration"`
}

type AddStepReq struct {
	Description string `json:"description" binding:"required"`
	Duration    int64  `json:"duration" binding:"required"`
}
