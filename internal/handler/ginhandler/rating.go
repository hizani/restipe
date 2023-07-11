package ginhandler

import (
	"net/http"
	"restipe/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary rate recipe
// @Security ApiKeyAuth
// @Tags recipe rating
// @Description  rate recipe from 1 to 5
// @ID rate-recipe
// @Accept json
// @Produce json
// @Param input body model.RateReq true "rate number"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/rates [post]
func (h *GinHandler) RateRecipe(c *gin.Context) {
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

	var input model.RateReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Recipe.RateRecipe(userId, recipeId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// @Summary rerate recipe
// @Security ApiKeyAuth
// @Tags recipe rating
// @Description  rerate recipe from 1 to 5
// @ID rate-recipe
// @Accept json
// @Produce json
// @Param input body model.RateReq true "rate number"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/rates [put]
func (h *GinHandler) RerateRecipe(c *gin.Context) {
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

	var input model.RateReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Recipe.RerateRecipe(userId, recipeId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary rate recipe
// @Security ApiKeyAuth
// @Tags recipe rating
// @Description  rate recipe from 1 to 5
// @ID rate-recipe
// @Accept json
// @Produce json
// @Param input body model.RateReq true "rate number"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/rates [post]
func (h *GinHandler) UploadImage(c *gin.Context) {
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

	var input model.RateReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Recipe.RateRecipe(userId, recipeId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}
