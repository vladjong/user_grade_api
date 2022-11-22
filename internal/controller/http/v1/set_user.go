package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *RouterTwo) SetUser(c *gin.Context) {
	c.Status(http.StatusOK)
}
