package app

import (
	"context"
	"time"

	service "github.com/boris-on/hflabs-test-task/internal/service/usecase"
	"github.com/sirupsen/logrus"
)

type HFLApp interface {
	Run()
}

type hflApp struct {
	HFLAppConfig
	ConfluenceService service.ConfluenceService
	GdocsService      service.GdocsService
}

type HFLAppConfig struct {
	Ctx               context.Context
	TableUpdatePeriod int //hours
	ConfluenceUrl     string
	GdocsID           string
	Logger            *logrus.Logger
	ClientConfig      string
}

func NewHFLApp(config HFLAppConfig) HFLApp {
	gdocsService, err := service.NewGdocsService(config.GdocsID, config.ClientConfig)
	if err != nil {
		config.Logger.Fatalf("error while initializng client: %s", err.Error())
	}
	return &hflApp{
		HFLAppConfig:      config,
		ConfluenceService: service.NewConfluenceService(config.ConfluenceUrl),
		GdocsService:      gdocsService,
	}
}

func ParseAndUpdate(h hflApp) {
	codeStatuses, err := h.ConfluenceService.GetCodes()
	if err != nil {
		h.Logger.Errorf("error parsing codes: %s", err.Error())
	}
	h.Logger.Info("codes were parsed")

	err = h.GdocsService.UpdateTable(codeStatuses)
	if err != nil {
		h.Logger.Errorf("error updating table: %s", err.Error())
	}
	h.Logger.Info("table was updated")
}

func (h hflApp) Run() {

	tableUpdateTicker := time.NewTicker(time.Duration(h.TableUpdatePeriod) * time.Hour)
	ParseAndUpdate(h)

	for {
		<-tableUpdateTicker.C
		go ParseAndUpdate(h)
	}
}
