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
	"github.com/vladjong/user_grade_api/pkg/server"
)

type App struct {
}

func (a *App) Run() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		a.startHTTP(&v1.RouterOne{}, viper.GetString("port_one"), &wg)
	}()
	wg.Add(1)
	go func() {
		a.startHTTP(&v1.RouterTwo{}, viper.GetString("port_two"), &wg)
	}()
	wg.Wait()
}

func (a *App) startHTTP(router v1.Router, port string, wg *sync.WaitGroup) {
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
