package model

type Ingredient struct {
	Id       int    `json:"id" binding:"required" db:"id"`
	Name     string `json:"name" db:"name"`
	Quantity int    `json:"quantity" binding:"required" db:"quantity"`
}

type AddIngredientReq struct {
	IngredientId int `json:"ingredient_id" binding:"required"`
	Quantity     int `json:"quantity" binding:"required"`
}
