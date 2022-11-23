package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHeader() []string {
	return []string{"user_id", "postpaid_limit", "spp", "shipping_fee", "return_fee"}
}

func (r *RouterOne) getBackup(c *gin.Context) {
	// запрос к бд массив структур
	// fileWoker := fileworker.WorkerCsv{}
	// fileWoker.Record(getHeader())
	users, err := r.Storage.GetBackup()
	if err != nil {
		ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}
