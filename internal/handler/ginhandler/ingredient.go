package ginhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *GinHandler) addIngredientToRecipe(c *gin.Context) {

}

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
