package v1

import "github.com/gin-gonic/gin"

type Router interface {
	NewRouter(handler *gin.Engine)
}
