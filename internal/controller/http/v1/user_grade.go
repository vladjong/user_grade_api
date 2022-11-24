package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vladjong/user_grade_api/internal/entity"
)

func (r *RouterOne) getUser(c *gin.Context) {
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

func (r *RouterTwo) setUser(c *gin.Context) {
	var inputUser entity.UserGrade
	if err := c.BindJSON(&inputUser); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if _, err := strconv.Atoi(inputUser.UserId); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := r.Producer.SendMessage(inputUser); err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}
