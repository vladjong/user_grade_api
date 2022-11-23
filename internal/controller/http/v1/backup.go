package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getHeader() []string {
	return []string{"user_id", "postpaid_limit", "spp", "shipping_fee", "return_fee", time.Now().String()}
}

func (r *RouterOne) getBackup(c *gin.Context) {
	users, err := r.Storage.GetBackup()
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	filename, err := r.FileWorkerer.Record(users, getHeader())
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Filename": filename,
	})
}
