package ginhandler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"restipe/internal/common"
	"restipe/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// @Summary upload recipe image
// @Security ApiKeyAuth
// @Tags recipe image
// @Description  upload recipe image. multipart key should be named as "image"
// @ID upload-recipe-img
// @Accept mpfd
// @Produce json
// @Param input formData file true "uploaded image"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/images [post]
func (h *GinHandler) uploadRecipeImg(c *gin.Context) {
	f, err := c.FormFile("image")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("can't get image: %s", err))
		return
	}

	if f == nil {
		newErrorResponse(c, http.StatusBadRequest, "no image provided")
		return
	}

	filename := f.Filename
	ext := filepath.Ext(filename)
	if ext != ".jpeg" && ext != ".jpg" && ext != ".png" {
		newErrorResponse(c, http.StatusBadRequest, "image should be .jpg or .png")
	}

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

	fileId := uuid.New()
	f.Filename = fileId.String() + ext
	oldFile, err := h.service.Recipe.UpdateRecipeImgFilename(userId, recipeId, &f.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	imgPath := fmt.Sprintf("%s%c%s%c%d%c%s",
		common.ImagesPath, os.PathSeparator, "recipe",
		os.PathSeparator, recipeId, os.PathSeparator, f.Filename,
	)
	if err := c.SaveUploadedFile(f, imgPath); err != nil {
		_, _ = h.service.Recipe.UpdateRecipeImgFilename(userId, recipeId, nil)
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if oldFile != nil && *oldFile != "" {
		if err := os.Remove(filepath.Dir(imgPath) + string(os.PathSeparator) + *oldFile); err != nil {
			log.Printf("can't remove old image: %s", err.Error())
		}
	}
	c.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary download recipe images
// @Tags recipe image
// @Description download an image of a recipe
// @ID download-recipe-img
// @Produce json
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/image [get]
func (h *GinHandler) downloadRecipeImg(c *gin.Context) {
	recipeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	filename, err := h.service.Recipe.GetRecipeImgFilename(recipeId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.File(fmt.Sprintf("%s%c%s%c%d%c%s",
		common.ImagesPath, os.PathSeparator, "recipe",
		os.PathSeparator, recipeId, os.PathSeparator, *filename,
	))
}
