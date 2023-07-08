package model

type Ingredient struct {
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity" binding:"required"`
}
