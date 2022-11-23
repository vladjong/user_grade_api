package app

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	v1 "github.com/vladjong/user_grade_api/internal/controller/http/v1"
	"github.com/vladjong/user_grade_api/internal/storage"
	jsondb "github.com/vladjong/user_grade_api/internal/storage/json_db"
	"github.com/vladjong/user_grade_api/pkg/kafka/producer"
	"github.com/vladjong/user_grade_api/pkg/server"
)

type app struct {
	storage  storage.UserStorager
	producer producer.Producer
}

func New() *app {
	storage := jsondb.New()

	options := producer.ProducerOptions{
		KafkaTopic:  viper.GetString("kafka_topic"),
		BrokersList: []string{viper.GetString("brokers_list")},
	}
	producer, err := producer.New(&options)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return &app{
		storage:  storage,
		producer: producer,
	}
}

func (a *app) Run() {
	var wg sync.WaitGroup
	routerOne := v1.RouterOne{
		Storage: a.storage,
	}
	routerTwo := v1.RouterTwo{
		Storage:  a.storage,
		Producer: a.producer,
	}
	// options := producer.ProducerOptions{
	// 	KafkaTopic:  viper.GetString("kafka_topic"),
	// 	BrokersList: []string{viper.GetString("brokers_list")},
	// }
	// producer, err := producer.New(&options)
	// if err != nil {
	// 	logrus.Fatal(err.Error())
	// }

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	producer, err := producer.New(&options)
	// 	if err != nil {
	// 		logrus.Fatal(err.Error())
	// 	}
	// }()

	wg.Add(1)
	go func() {
		a.startHTTP(&routerOne, viper.GetString("port_one"), &wg)
	}()

	wg.Add(1)
	go func() {
		a.startHTTP(&routerTwo, viper.GetString("port_two"), &wg)
	}()

	wg.Wait()
}

func (a *app) startHTTP(router v1.Router, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	logrus.Info("HTTP server initializing")
	server := new(server.Server)
	handler := gin.Default()
	router.NewRouter(handler)
	go func() {
		if err := server.Run(port, handler); err != nil {
			logrus.Fatalf("occured while running HTTP server: %s", err.Error())
		}
	}()
	logrus.Info("HTTP server start")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Info("HTTP Server Shutdown")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("occured on server shutdown: %s", err.Error())
	}
}
