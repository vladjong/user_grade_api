package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vladjong/user_grade_api/internal/entity"
	"github.com/vladjong/user_grade_api/internal/storage"
	"github.com/vladjong/user_grade_api/pkg/kafka/producer"
)

type RouterTwo struct {
	Storage  storage.UserStorager
	Producer producer.Producer
}

func (r *RouterTwo) NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	api := handler.Group("/api", basicAuth)
	{
		api.POST("/", r.setUser)
	}
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
	if err := r.Storage.SetUser(inputUser); err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	r.Producer.SendMessage(inputUser)
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}
