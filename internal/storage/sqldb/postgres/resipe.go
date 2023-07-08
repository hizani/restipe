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

func (r *RecipeStorage) Create(userId int, recipe model.CreateRecipeReq) (int, error) {
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

func (r *RecipeStorage) GetAll(recipe model.GetAllRecipesReq) ([]model.Recipe, error) {
	recipes := []model.Recipe{}
	query := strings.Builder{}
	query.Grow(256)
	query.WriteString("SELECT rd.id, rd.name, rd.description, rd.author  FROM ")
	recipeQuery := fmt.Sprintf("SELECT * FROM %s", recipeTable)
	groupBy := "GROUP BY rd.id, rd.name, rd.description, rd.author "

	durationSort := strings.ToUpper(recipe.DurationSort)
	if durationSort == "ASC" || durationSort == "DESC" {
		recipeQuery = fmt.Sprintf("SELECT r.*, SUM(st.duration) dur FROM %s r "+
			"JOIN %s st on st.recipe_id  = r.id GROUP BY r.id",
			recipeTable, stepTable,
		)
		groupBy = "GROUP BY rd.id, rd.name, rd.description, rd.author, rd.dur "
	}
	query.WriteString(fmt.Sprintf("(%s) rd ", recipeQuery))

	havingIngredients := ""
	if len(recipe.IngredientFilter) != 0 {
		query.WriteString(
			fmt.Sprintf("JOIN %s ir ON rd.id = ir.recipe_id AND ir.ingredient_id in (",
				ingredientRecipeTable,
			),
		)
		for i, v := range recipe.IngredientFilter {
			query.WriteString(fmt.Sprintf("%d", v))
			if i != len(recipe.IngredientFilter)-1 {
				query.WriteString(",")
			}
		}
		query.WriteString(") ")
		havingIngredients = fmt.Sprintf("HAVING COUNT(ir.ingredient_id) = %d ", len(recipe.IngredientFilter))
	}
	if recipe.Author != 0 {
		query.WriteString(fmt.Sprintf("WHERE rd.author = %d ", recipe.Author))
	}

	query.WriteString(groupBy)
	query.WriteString(havingIngredients)
	if durationSort == "ASC" || durationSort == "DESC" {
		query.WriteString(fmt.Sprintf("ORDER BY rd.dur %s", durationSort))
	}

	err := r.db.Select(&recipes, query.String())

	return recipes, err
}

func (r *RecipeStorage) GetById(recipeId int) (model.Recipe, error) {
	var recipe model.Recipe
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", recipeTable)
	err := r.db.Get(&recipe, query, recipeId)
	return recipe, err
}
