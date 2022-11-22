package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vladjong/user_grade_api/internal/entity"
)

func (r *RouterTwo) SetUser(c *gin.Context) {
	var inputUser entity.UserGrade
	if err := c.BindJSON(&inputUser); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, err := strconv.Atoi(inputUser.UserId); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := r.Storage.SetUser(inputUser); err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}
