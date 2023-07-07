package postgres

import (
	"fmt"
	"restipe/internal/model"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type RecipeStorage struct {
	db *sqlx.DB
}

func NewRecipeStorage(db *sqlx.DB) *RecipeStorage {
	return &RecipeStorage{db}
}

func (r *RecipeStorage) Create(userId int, recipe model.CreateRecipe) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createRecipeQuery := fmt.Sprintf(
		"INSERT INTO %s (name, description, author) VALUES ($1, $2, $3) RETURNING id", recipeTable,
	)
	row := tx.QueryRow(createRecipeQuery, recipe.Name, recipe.Description, userId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	if len(recipe.Ingredients) != 0 {
		createRecipeIngredients := strings.Builder{}
		if _, err := createRecipeIngredients.WriteString(
			fmt.Sprintf("INSERT INTO %s (recipe_id, ingredient_id, quantity) VALUES ", ingredientRecipeTable),
		); err != nil {
			tx.Rollback()
			return 0, err
		}
		for i, v := range recipe.Ingredients {
			createRecipeIngredients.WriteString(
				fmt.Sprintf("(%d, %d, %d)", id, v.Id, v.Quantity),
			)
			if i < len(recipe.Ingredients)-1 {
				createRecipeIngredients.WriteString(", ")
			}
		}
		if _, err := tx.Exec(createRecipeIngredients.String()); err != nil {
			tx.Rollback()
			return 0, err
		}

	}

	if len(recipe.Steps) != 0 {
		createSteps := strings.Builder{}
		if _, err := createSteps.WriteString(
			fmt.Sprintf("INSERT INTO %s (recipe_id, number, description, duration) VALUES ", stepTable),
		); err != nil {
			tx.Rollback()
			return 0, err
		}
		for i, v := range recipe.Steps {

			createSteps.WriteString(
				fmt.Sprintf("(%d, %d, '%s', '%ds')", id, i+1, v.Description, int64(time.Duration(v.Duration).Seconds())),
			)
			if i < len(recipe.Steps)-1 {
				createSteps.WriteString(", ")
			}
		}
		if _, err := tx.Exec(createSteps.String()); err != nil {
			fmt.Println(createSteps.String())
			tx.Rollback()
			return 0, err
		}

	}

	return id, tx.Commit()
}
