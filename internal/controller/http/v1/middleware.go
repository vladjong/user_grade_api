package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func basicAuth(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == "admin" && password == "qwerty" {
		logrus.Info("user authenticated")
	} else {
		ErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}
}
