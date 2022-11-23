package v1

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func basicAuth(c *gin.Context) {
	user, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && user == os.Getenv("LOGIN") && password == os.Getenv("PASSWORD") {
		logrus.Info("user authenticated")
	} else {
		ErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}
}
