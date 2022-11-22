package main

import (
	"github.com/sirupsen/logrus"
	"github.com/vladjong/user_grade_api/config"
	"github.com/vladjong/user_grade_api/internal/app"
)

func main() {
	logrus.Info("init config")
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}
	app := app.App{}
	app.Run()
}
