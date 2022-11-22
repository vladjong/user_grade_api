package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (r *RouterOne) GetUser(c *gin.Context) {
	id := c.Param("id")
	if _, err := strconv.Atoi(id); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := r.Storage.GetUser(id)
	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
