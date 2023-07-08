package storage

import (
	"restipe/internal/model"
	"restipe/internal/storage/sqldb"
	"restipe/internal/storage/sqldb/postgres"
)

type Authorization interface {
	SignupUser(user model.SignupUserReq) (int, error)
	SigninUser(user model.SigninUserReq) (int, error)
}

type Recipe interface {
	Create(userId int, recipe model.CreateRecipeReq) (int, error)
	GetAll(recipe model.GetAllRecipesReq) ([]model.Recipe, error)
}

type Storage struct {
	Authorization
	Recipe
}

func NewPostgres(cfg sqldb.Config) (*Storage, error) {
	db, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}
	return &Storage{
		Authorization: postgres.NewAuthStoarge(db),
		Recipe:        postgres.NewRecipeStorage(db),
	}, err
}
