package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/boris-on/hflabs-test-task/internal/app"
	"github.com/boris-on/hflabs-test-task/internal/service/logger"
	"github.com/boris-on/hflabs-test-task/internal/service/pprofing"

	"github.com/spf13/viper"
)

const (
	TableUpdatePeriod = 12 //hours
)

func main() {
	log := logger.New()
	ctx := context.Background()

	if err := InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	go pprofing.InitPprofing(ctx, log, viper.GetString("PprofPort"))

	appConfig := app.HFLAppConfig{
		Ctx:               ctx,
		TableUpdatePeriod: TableUpdatePeriod,
		ConfluenceUrl:     viper.GetString("ConfluenceUrl"),
		GdocsID:           viper.GetString("GdocsID"),
		ClientConfig:      viper.GetString("ClientConfig"),
		Logger:            log,
	}

	srv := app.NewHFLApp(appConfig)

	log.Info("starting app")
	go srv.Run()

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)
	<-gracefulShutdown

	log.Info("app shutting down")
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
