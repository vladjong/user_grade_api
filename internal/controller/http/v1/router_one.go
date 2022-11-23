package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vladjong/user_grade_api/internal/storage"
	"github.com/vladjong/user_grade_api/pkg/fileworker"
)

type RouterOne struct {
	Storage      storage.UserStorager
	FileWorkerer fileworker.FileWorkerer
}

func (r *RouterOne) NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	api := handler.Group("/api")
	{
		api.GET("/:id", r.getUser)
		api.GET("/backup", r.getBackup)
	}
}

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
