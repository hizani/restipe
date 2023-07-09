package ginhandler

import (
	"net/http"
	"restipe/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary add ingredient
// @Security ApiKeyAuth
// @Tags recipe ingredient
// @Description  add ingredient to recipe
// @ID add-ingredient
// @Accept json
// @Produce json
// @Param input body model.AddIngredientReq true "ingredient info"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/ingredients [post]
func (h *GinHandler) addIngredientToRecipe(c *gin.Context) {
	userId := c.GetInt(userCtx)
	if userId == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	recipeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input model.AddIngredientReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Recipe.AddIngredientToRecipe(userId, recipeId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// @Summary get ingredients
// @Tags recipe ingredient
// @Description  get all recipe ingredients
// @ID get-ingredients
// @Produce json
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/ingredients [get]
func (h *GinHandler) getAllIngredientsFromRecipe(c *gin.Context) {
	recipeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	ingredients, err := h.service.Recipe.GetAllIngredientsFromRecipe(recipeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, ingredients)
}

// @Summary remove ingredient
// @Security ApiKeyAuth
// @Tags recipe ingredient
// @Description  remove ingredient from recipe
// @ID remove-ingredient
// @Produce json
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/ingredients/{ingid} [delete]
func (h *GinHandler) removeIngredientFromRecipe(c *gin.Context) {
	userId := c.GetInt(userCtx)
	if userId == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	recipeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	ingredientId, err := strconv.Atoi(c.Param("ingid"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Recipe.RemoveIngredientFromRecipe(userId, recipeId, ingredientId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{})
}
