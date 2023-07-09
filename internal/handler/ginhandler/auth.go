package ginhandler

import (
	"net/http"
	"restipe/internal/model"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary SingUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body model.SignupUserReq true "account info"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /auth/signup [post]
func (h *GinHandler) signUp(c *gin.Context) {
	var input model.SignupUserReq

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.service.Authorization.SignupUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})

}

// @Summary SingIn
// @Tags auth
// @Description log into user account
// @ID login
// @Accept json
// @Produce json
// @Param input body model.SigninUserReq true "account info"
// @Success 200  {integer} integer 1
// @Failure 400,404 {object} ginHandlerError
// @Failure 500 {object} ginHandlerError
// @failure default {object} ginHandlerError
// @Router /auth/signin [post]
func (h *GinHandler) signIn(c *gin.Context) {
	var input model.SigninUserReq

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.service.Authorization.SigninUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}

func (h *GinHandler) Authorize(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if headerParts[0] != "Bear" && len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.service.Authorization.Authorize(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}
