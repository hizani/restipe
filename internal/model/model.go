package model

type User struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Login    string `json:"Login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignupUser struct {
	Name     string `json:"name" binding:"required"`
	Login    string `json:"Login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SigninUser struct {
	Login    string `json:"Login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateRecipe struct {
	Id          int                `json:"id"`
	Name        string             `json:"name" binding:"required"`
	Description string             `json:"description"`
	Ingredients []InsertIngredient `json:"ingredients" binding:"required"`
	Steps       []CreateStep       `json:"steps" binding:"required"`
}

type InsertIngredient struct {
	Id       int `json:"id" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

type CreateStep struct {
	Id          int    `json:"id"`
	Description string `json:"description" binding:"required"`
	Duration    int64  `json:"duration" binding:"required"`
}
