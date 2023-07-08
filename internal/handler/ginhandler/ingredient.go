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

}
