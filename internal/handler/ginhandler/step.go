package ginhandler

import (
	"net/http"
	"restipe/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *GinHandler) addStepToRecipe(c *gin.Context) {
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

	var input model.AddStepReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Recipe.AddStepToRecipe(userId, recipeId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *GinHandler) getAllStepsFromRecipe(c *gin.Context) {
	recipeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	ingredients, err := h.service.Recipe.GetAllStepsFromRecipe(recipeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, ingredients)
}

func (h *GinHandler) deleteStep(c *gin.Context) {}
