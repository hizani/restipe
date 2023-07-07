package ginhandler

import "github.com/gin-gonic/gin"

type ginHandlerError struct {
	Message string `json:"message"`
}

func (ge *ginHandlerError) Error() string {
	return ge.Message
}

func newErrorResponce(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ginHandlerError{message})
}
