package model

import "time"

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

type Recipe struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	AuthorId    int          `json:"author"`
	Ingredients []Ingredient `json:"ingredients"`
	Steps       []Step       `json:"steps"`
}

type Ingredient struct {
	Id       int    `json:"id"`
	RecipeId int    `json:"recipe_id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Step struct {
	Id          int           `json:"id"`
	RecipeId    int           `json:"recipe_id"`
	Number      int           `json:"number"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"`
}
