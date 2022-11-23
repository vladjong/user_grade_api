package v1

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/vladjong/user_grade_api/internal/entity"
	"github.com/vladjong/user_grade_api/internal/storage"
	"github.com/vladjong/user_grade_api/pkg/kafka/consumer"
	"github.com/vladjong/user_grade_api/pkg/kafka/producer"
)

type RouterTwo struct {
	Storage  storage.UserStorager
	Producer producer.Producer
	Consumer consumer.Consumer
}

func (r *RouterTwo) NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	api := handler.Group("/api", basicAuth)
	{
		api.POST("/", r.setUser)
		api.GET("/backup", r.getBackup)
	}

	go func() {
		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()
		if err := consumer.StartConsumerGroup(ctx, consumer.ConsumerOptions{
			KafkaTopic:         viper.GetString("kafka_topic"),
			KafkaConsumerGroup: viper.GetString("kafka_consumer_group"),
			BrokersList:        []string{viper.GetString("brokers_list")},
			Assignor:           viper.GetString("assignor"),
		}, r.Storage); err != nil {
			logrus.Fatal(err)
		}
		<-ctx.Done()
	}()
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
	r.Producer.SendMessage(inputUser)
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}

func (r *RouterTwo) getBackup(c *gin.Context) {
	// var inputUser entity.UserGrade
	// if err := c.BindJSON(&inputUser); err != nil {
	// 	ErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// if _, err := strconv.Atoi(inputUser.UserId); err != nil {
	// 	ErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// r.Producer.SendMessage(inputUser)
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "Ok",
	})
}
