package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"restipe/internal/model"
	"strings"

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
				fmt.Sprintf("(%d, %d, '%s', '%ds')", id, i+1, v.Description, v.Duration),
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

func (r *RecipeStorage) GetAllIngredientsFromRecipe(recipeId int) ([]model.Ingredient, error) {
	var ingredients []model.Ingredient

	query := fmt.Sprintf("SELECT i.*, ir.quantity FROM %s i "+
		"JOIN %s ir ON ir.ingredient_id = i.id WHERE ir.recipe_id = $1", ingredientTable, ingredientRecipeTable)
	err := r.db.Select(&ingredients, query, recipeId)
	return ingredients, err
}

func (r *RecipeStorage) GetAllStepsFromRecipe(recipeId int) ([]model.Step, error) {
	var steps []model.Step

	query := fmt.Sprintf("SELECT s.id, s.number, s.description, EXTRACT(EPOCH FROM s.duration)::int duration "+
		"FROM %s s WHERE s.recipe_id = $1", stepTable)
	err := r.db.Select(&steps, query, recipeId)
	return steps, err
}

func (r *RecipeStorage) AddStepToRecipe(userId int, recipeId int, step model.AddStepReq) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int

	var author int
	SelectRecipeAuthor := fmt.Sprintf(
		"SELECT r.author FROM %s u JOIN %s r ON u.id = r.author WHERE r.id = $1 AND u.id = $2",
		userTable, recipeTable,
	)
	row := tx.QueryRow(SelectRecipeAuthor, recipeId, userId)
	if err := row.Scan(&author); err != nil {
		if err != sql.ErrNoRows {
			tx.Rollback()
			return 0, err
		}
		return 0, errors.New("wrong author")
	}

	var maxNumber int
	SelectMaxNumQuery := fmt.Sprintf(
		"SELECT number FROM %s WHERE recipe_id = $1 ORDER BY number DESC LIMIT 1", stepTable,
	)
	row = tx.QueryRow(SelectMaxNumQuery, recipeId)
	if err := row.Scan(&maxNumber); err != nil {
		if err != sql.ErrNoRows {
			tx.Rollback()
			return 0, err
		}
	}

	InsertStepQuery := fmt.Sprintf(
		"INSERT INTO %s (recipe_id, number, description, duration) VALUES ($1, $2, $3, $4) RETURNING id",
		stepTable,
	)
	row = tx.QueryRow(InsertStepQuery, recipeId, maxNumber+1, step.Description, fmt.Sprintf("%ds", step.Duration))
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *RecipeStorage) RemoveStepFromRecipe(userId, recipeId, stepId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	var author int
	SelectRecipeAuthor := fmt.Sprintf(
		"SELECT r.author FROM %s u JOIN %s r ON u.id = r.author WHERE r.id = $1 AND u.id = $2",
		userTable, recipeTable,
	)
	row := tx.QueryRow(SelectRecipeAuthor, recipeId, userId)
	if err := row.Scan(&author); err != nil {
		if err != sql.ErrNoRows {
			tx.Rollback()
			return err
		}
		return errors.New("wrong author")
	}

	var number int
	GetRemovedNumber := fmt.Sprintf(
		"SELECT number FROM %s WHERE recipe_id = $1 AND id = $2",
		stepTable,
	)
	row = tx.QueryRow(GetRemovedNumber, recipeId, stepId)
	if err := row.Scan(&number); err != nil {
		tx.Rollback()
		return err
	}

	UpdateStepsNumber := fmt.Sprintf(
		"UPDATE %s SET number = number - 1 WHERE number > $1",
		stepTable,
	)

	if _, err := tx.Exec(UpdateStepsNumber, number); err != nil {
		tx.Rollback()
		return err
	}

	RemoveStepQuery := fmt.Sprintf(
		"DELETE FROM %s WHERE recipe_id = $1 AND id = $2",
		stepTable,
	)
	if _, err := tx.Exec(RemoveStepQuery, recipeId, stepId); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}
