package postgres

import (
	"fmt"
	"restipe/internal/storage/sqldb"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	userTable             = "\"user\""
	recipeTable           = "\"recipe\""
	ingredientTable       = "\"ingredient\""
	ingredientRecipeTable = "\"ingredient_recipe\""
	stepTable             = "\"step\""
	ratingTable           = "\"rating\""
)

func New(cfg sqldb.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.DBName, cfg.Username, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
