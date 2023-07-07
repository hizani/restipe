package ginhandler

import "github.com/gin-gonic/gin"

type ginError struct {
	Message string `json:"message"`
}

func (ge *ginError) Error() string {
	return ge.Message
}

func newErrorResponce(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, ginError{message})
}
