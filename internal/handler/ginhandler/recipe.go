package ginhandler

import (
	"net/http"
	"restipe/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *GinHandler) createRecipe(c *gin.Context) {
	userId := c.GetInt(userCtx)
	if userId == 0 {
		newErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input model.CreateRecipeReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Recipe.Create(userId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *GinHandler) getAllRecipes(c *gin.Context) {
	var input model.GetAllRecipesReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}
	recipes, err := h.service.Recipe.GetAll(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, recipes)

}

func (h *GinHandler) getRecipeById(c *gin.Context) {

}

func (h *GinHandler) updateRecipe(c *gin.Context) {

}

func (h *GinHandler) deleteRecipe(c *gin.Context) {

}

func (h *GinHandler) getAllUserRecipes(c *gin.Context) {

}

func (h *GinHandler) getRecipesWithIngredients(c *gin.Context) {}
