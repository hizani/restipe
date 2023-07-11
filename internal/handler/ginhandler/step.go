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

// @Summary add step
// @Security ApiKeyAuth
// @Tags recipe step
// @Description  add step to recipe
// @ID add-step
// @Accept json
// @Produce json
// @Param input body model.AddStepReq true "ingredient info"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/step [post]
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

// @Summary get steps
// @Tags recipe step
// @Description  get all recipe steps
// @ID get-steps
// @Produce json
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/steps [get]
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

// @Summary remove step
// @Security ApiKeyAuth
// @Tags recipe step
// @Description  remove step from recipe
// @ID remove-step
// @Produce json
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/ingredients/{number} [delete]
func (h *GinHandler) removeStepFromRecipe(c *gin.Context) {
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

	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.Recipe.RemoveStepFromRecipe(userId, recipeId, number)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{})
}

// @Summary upload step image
// @Security ApiKeyAuth
// @Tags recipe image
// @Description  upload step image. multipart key should be named as "image"
// @ID upload-step-img
// @Accept mpfd
// @Produce json
// @Param input formData file true "uploaded image"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/steps/{number}/images [post]
func (h *GinHandler) uploadStepImg(c *gin.Context) {
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

	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fileId := uuid.New().String()
	f.Filename = fileId + ext
	oldFile, err := h.service.Recipe.UpdateStepImgFilename(userId, recipeId, number, &f.Filename)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	imgPath := fmt.Sprintf("%s%c%s%c%d%c%s%c%d%c%s",
		common.ImagesPath, os.PathSeparator, "recipe",
		os.PathSeparator, recipeId, os.PathSeparator, "step",
		os.PathSeparator, number, os.PathSeparator, f.Filename,
	)
	if err := c.SaveUploadedFile(f, imgPath); err != nil {
		_, _ = h.service.Recipe.UpdateStepImgFilename(userId, recipeId, number, nil)
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

// @Summary download step image
// @Tags recipe image
// @Description download an image of a recipe
// @ID download-step-img
// @Produce json
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /api/recipes/{id}/steps/{number}/images [get]
func (h *GinHandler) downloadStepImg(c *gin.Context) {
	recipeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	filename, err := h.service.Recipe.GetStepImgFilename(recipeId, number)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(*filename)
	c.File(fmt.Sprintf("%s%c%s%c%d%c%s%c%d%c%s",
		common.ImagesPath, os.PathSeparator, "recipe",
		os.PathSeparator, recipeId, os.PathSeparator, "step",
		os.PathSeparator, number, os.PathSeparator, *filename,
	))
}
