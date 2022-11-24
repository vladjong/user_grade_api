package v1

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/user_grade_api/internal/entity"
)

const gzipFilename = "data/backup.csv.gz"

func getHeader() []string {
	return []string{"user_id", "postpaid_limit", "spp", "shipping_fee", "return_fee"}
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
	inputBackup := entity.Backup{Filename: filename}
	c.JSON(http.StatusOK, inputBackup)
}

func (r *RouterTwo) readBackup(wg *sync.WaitGroup) {
	wg.Done()
	logrus.Info("start read backup")
	inputBackup := entity.Backup{Filename: gzipFilename}
	users, err := r.FileWorkerer.GetRecord(inputBackup.Filename)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	if err := r.Storage.SetBackup(users); err != nil {
		logrus.Error(err.Error())
		return
	}
}
