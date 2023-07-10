package postgres

import (
	"database/sql"
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
				fmt.Sprintf("(%d, %d, %d)", id, v.IngredientId, v.Quantity),
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

func (r *RecipeStorage) Delete(userId, recipeId int) error {
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
		return nil
	}

	DeleteRecipeQuery := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1",
		recipeTable,
	)

	if _, err = tx.Exec(DeleteRecipeQuery, recipeId); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *RecipeStorage) GetAll(recipe model.GetAllRecipesReq) ([]model.AllRecipeResp, error) {
	setQuery := make([]string, 7)

	if len(recipe.IngredientFilter) > 0 {
		setQuery[0] =
			fmt.Sprintf("JOIN %s ir ON rd.id = ir.recipe_id AND ir.ingredient_id in (%s)",
				ingredientRecipeTable,
				strings.Trim(strings.Join(strings.Fields(fmt.Sprint(recipe.IngredientFilter)), ","), "[]"),
			)
		setQuery[4] = fmt.Sprintf("HAVING COUNT(ir.ingredient_id) = %d ", len(recipe.IngredientFilter))
	}

	whereStart := "WHERE"
	if recipe.DurationFilter != nil {
		duration := fmt.Sprintf("'%ds'", *recipe.DurationFilter)
		setQuery[1] =
			fmt.Sprintf("%s rd.dur = %s",
				whereStart,
				duration,
			)
		whereStart = "AND"
	}

	if recipe.Author != nil {
		setQuery[2] = fmt.Sprintf("%s rd.author = %d", whereStart, *recipe.Author)
	}

	if recipe.RatingFilter != nil {
		setQuery[3] = fmt.Sprintf("%s rt.avg_rating = %f", whereStart, *recipe.RatingFilter)
	}

	orderStart := "ORDER BY"
	if recipe.DurationSort != nil {
		durSort := strings.ToUpper(*recipe.DurationSort)
		if durSort == "ASC" || durSort == "DESC" {
			setQuery[5] = fmt.Sprintf("%s rd.dur %s", orderStart, durSort)
			orderStart = ","
		}
	}
	if recipe.RatingSort != nil {
		ratingSort := strings.ToUpper(*recipe.RatingSort)
		if ratingSort == "ASC" || ratingSort == "DESC" {
			setQuery[6] = fmt.Sprintf("%s avg_rating %s", orderStart, ratingSort)
		}
	}

	query := fmt.Sprintf("SELECT rd.id, rd.name, rd.description, rd.author, "+
		"EXTRACT(EPOCH FROM rd.dur)::int duration, COALESCE(rt.avg_rating, 0) avg_rating FROM "+
		"(SElECT r.*, SUM(st.duration) dur FROM %s r JOIN %s st ON st.recipe_id = r.id GROUP BY r.id) rd "+
		"LEFT JOIN (SElECT rt.recipe_id, AVG(rt.rating)::numeric(2,1) avg_rating FROM %s rt GROUP BY rt.recipe_id) rt "+
		"ON rt.recipe_id = rd.id "+
		"%s %s %s %s GROUP BY rd.id, rd.name, rd.description, rd.author, rd.dur, avg_rating %s %s%s",
		recipeTable, stepTable, ratingTable,
		setQuery[0], setQuery[1], setQuery[2], setQuery[3], setQuery[4], setQuery[5], setQuery[6])

	fmt.Println(query)

	recipes := []model.AllRecipeResp{}
	err := r.db.Select(&recipes, query)

	return recipes, err
}

func (r *RecipeStorage) Update(userId, recipeId int, recipe model.UpdateRecipeReq) error {
	setValues := make([]string, 0, 2)
	args := make([]interface{}, 0, 2)
	argId := 1

	if recipe.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *recipe.Name)
		argId++
	}

	if recipe.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *recipe.Description)
		argId++
	}

	if len(args) == 0 {
		return nil
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s r SET %s WHERE r.id = $%d AND r.author = $%d",
		recipeTable, setQuery, argId, argId+1)
	args = append(args, recipeId, userId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *RecipeStorage) GetById(recipeId int) (model.RecipeResp, error) {
	var recipe model.RecipeResp
	queryRecipe := fmt.Sprintf("SELECT r.*, COALESCE(AVG(rt.rating)::numeric(2, 1), 0) avg_rating FROM %s r "+
		"LEFT JOIN %s rt ON r.id = rt.recipe_id WHERE r.id = $1 GROUP BY r.id", recipeTable, ratingTable)
	err := r.db.Get(&recipe, queryRecipe, recipeId)
	if err != nil {
		return recipe, err
	}

	var ingredients []model.Ingredient
	queryIngredients := fmt.Sprintf("SELECT ingredient_id id, i.name, quantity FROM %s ir "+
		"JOIN %s i ON ir.ingredient_id = i.id WHERE recipe_id = $1", ingredientRecipeTable, ingredientTable)
	err = r.db.Select(&ingredients, queryIngredients, recipeId)
	if err != nil {
		return recipe, err
	}
	recipe.Ingredients = ingredients

	var steps []model.Step
	queryStep := fmt.Sprintf("SELECT id, number, description, EXTRACT(EPOCH FROM duration)::int duration "+
		"FROM %s WHERE recipe_id = $1", stepTable)
	err = r.db.Select(&steps, queryStep, recipeId)
	recipe.Steps = steps
	var durAcc int64
	for _, v := range steps {
		durAcc += v.Duration
	}
	recipe.Duration = durAcc

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
		return 0, nil
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

func (r *RecipeStorage) AddIngredientToRecipe(userId int, recipeId int, ingredient model.AddIngredientReq) (int, error) {
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
		return 0, nil
	}

	InsertIngredientQuery := fmt.Sprintf(
		"INSERT INTO %s (recipe_id, ingredient_id, quantity) VALUES ($1, $2, $3) RETURNING id",
		ingredientRecipeTable,
	)
	row = tx.QueryRow(InsertIngredientQuery, recipeId, ingredient.IngredientId, ingredient.Quantity)
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
		return nil
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

func (r *RecipeStorage) RemoveIngredientFromRecipe(userId, recipeId, ingredientId int) error {
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
		return nil
	}

	RemoveIngredientQuery := fmt.Sprintf(
		"DELETE FROM %s WHERE recipe_id = $1 AND ingredient_id = $2",
		ingredientRecipeTable,
	)
	if _, err := tx.Exec(RemoveIngredientQuery, recipeId, ingredientId); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}

func (r *RecipeStorage) RateRecipe(userId, recipeId int, rating model.RateReq) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (user_id, recipe_id, rating) VALUES ($1, $2, $3) RETURNING id",
		ratingTable,
	)
	row := r.db.QueryRow(query, userId, recipeId, rating.Rating)
	return id, row.Scan(&id)
}

func (r *RecipeStorage) RerateRecipe(userId, recipeId int, rating model.RateReq) error {
	query := fmt.Sprintf(
		"UPDATE %s SET rating = $1 WHERE user_id = $2 AND recipe_id = $3",
		ratingTable,
	)
	_, err := r.db.Exec(query, rating.Rating, userId, recipeId)
	return err
}

func (r *RecipeStorage) Close() error {
	return r.db.Close()
}
