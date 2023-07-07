package ginhandler

import (
	"net/http"
	"restipe/internal/model"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *GinHandler) signUp(c *gin.Context) {
	var input model.SignupUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.service.Authorization.SignupUser(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})

}

func (h *GinHandler) signIn(c *gin.Context) {
	var input model.SigninUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.service.Authorization.SigninUser(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func (h *GinHandler) Authorize(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponce(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if headerParts[0] != "Bear" && len(headerParts) != 2 {
		newErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.service.Authorization.Authorize(headerParts[1])
	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}
