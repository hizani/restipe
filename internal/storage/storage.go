package storage

import (
	"restipe/internal/model"
	"restipe/internal/storage/sqldb"
	"restipe/internal/storage/sqldb/postgres"
)

type Authorization interface {
	SignupUser(user model.SignupUser) (int, error)
	SigninUser(user model.SigninUser) (int, error)
}

type Recipe interface {
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
	}, err
}
