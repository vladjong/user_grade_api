package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouterTwo struct {
}

func (r *RouterTwo) NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	api := handler.Group("/v1")
	{
		api.GET("/:id", r.SetUser)
	}
}
