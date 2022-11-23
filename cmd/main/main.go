package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vladjong/user_grade_api/config"
	"github.com/vladjong/user_grade_api/internal/app"
)

func main() {
	logrus.Info("init config")
	if err := config.InitConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}
	logrus.Info("env variables initializing")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	app := app.New()
	app.Run()
}
