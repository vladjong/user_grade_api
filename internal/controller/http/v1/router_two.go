package v1

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/vladjong/user_grade_api/internal/storage"
	"github.com/vladjong/user_grade_api/pkg/fileworker"
	"github.com/vladjong/user_grade_api/pkg/kafka/consumer"
	"github.com/vladjong/user_grade_api/pkg/kafka/producer"
)

type RouterTwo struct {
	Storage      storage.UserStorager
	Producer     producer.Producer
	Consumer     consumer.Consumer
	FileWorkerer fileworker.FileWorkerer
}

func (r *RouterTwo) NewRouter(handler *gin.Engine) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	api := handler.Group("/api")
	{
		api.POST("/", basicAuth, r.setUser)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		r.readBackup(&wg)
	}()
	wg.Wait()

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
