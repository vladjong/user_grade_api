package v1

import "github.com/gin-gonic/gin"

type Response struct {
	Error string `json:"error"`
}

func ErrorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, Response{msg})
}
