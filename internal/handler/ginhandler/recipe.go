package ginhandler

import (
	"net/http"
	"restipe/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary create recipe
// @Security ApiKeyAuth
// @Tags recipe
// @Description create recipe
// @ID create-recipe
// @Accept json
// @Produce json
// @Param input body model.CreateRecipeReq true "recipe info"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes [post]
func (h *GinHandler) createRecipe(c *gin.Context) {
	userId := c.GetInt(userCtx)
	if userId == 0 {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input model.CreateRecipeReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Recipe.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// @Summary get all recipes
// @Tags recipe
// @Description get all recipes
// @ID get-all-recipes
// @Accept json
// @Produce json
// @Param input body model.GetAllRecipesReq true "recipe search request info"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes [get]
func (h *GinHandler) getAllRecipes(c *gin.Context) {
	var input model.GetAllRecipesReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	recipes, err := h.service.Recipe.GetAll(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, recipes)

}

// @Summary get recipe by id
// @Tags recipe
// @Description get recipe by id
// @ID get-recipe-by-id
// @Produce json
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id} [get]
func (h *GinHandler) getRecipeById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	recipe, err := h.service.Recipe.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, recipe)

}

// @Summary update recipe
// @Security ApiKeyAuth
// @Tags recipe
// @Description update recipe
// @ID update-recipe
// @Accept json
// @Produce json
// @Param input body model.UpdateRecipeReq true "recipe update info"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id} [put]
func (h *GinHandler) updateRecipe(c *gin.Context) {
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

	var input model.UpdateRecipeReq
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Recipe.Update(userId, recipeId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary delete recipe
// @Security ApiKeyAuth
// @Tags recipe
// @Description delete recipe
// @ID delete-recipe
// @Produce json
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id} [delete]
func (h *GinHandler) deleteRecipe(c *gin.Context) {
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

	err = h.service.Recipe.Delete(userId, recipeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{})
}
