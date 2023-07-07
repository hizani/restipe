package ginhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *GinHandler) createRecipe(c *gin.Context) {

}

func (h *GinHandler) getAllRecipe(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"aboba": "aboba",
	})

}

func (h *GinHandler) getRecipeById(c *gin.Context) {

}

func (h *GinHandler) updateRecipe(c *gin.Context) {

}

func (h *GinHandler) deleteRecipe(c *gin.Context) {

}
