package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *RouterOne) GetUser(c *gin.Context) {
	c.Status(http.StatusOK)
}
