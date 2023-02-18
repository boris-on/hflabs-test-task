package usecase

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/boris-on/hflabs-test-task/internal/service/models"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type GdocsService interface {
	UpdateTable([]models.StatusCode) error
}

type gdocsService struct {
	GdocsID     string
	GdocsClient *http.Client
}

func NewGdocsService(id, clientConfig string) (GdocsService, error) {
	data, err := ioutil.ReadFile(clientConfig)
	if err != nil {
		return &gdocsService{}, err
	}

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		return &gdocsService{}, err
	}

	client := conf.Client(context.TODO())
	return &gdocsService{GdocsID: id, GdocsClient: client}, nil
}

func (d gdocsService) UpdateTable(statusCodes []models.StatusCode) error {

	service := spreadsheet.NewServiceWithClient(d.GdocsClient)
	spreadsheet, err := service.FetchSpreadsheet(d.GdocsID)
	if err != nil {
		return err
	}

	sheet, err := spreadsheet.SheetByIndex(0)
	if err != nil {
		return err
	}

	sheet.Update(0, 0, "HTTP-код ответа")
	sheet.Update(0, 1, "Описание")
	for i, statusCode := range statusCodes {
		sheet.Update(i+1, 0, statusCode.Code)
		sheet.Update(i+1, 1, statusCode.Description)
	}

	err = sheet.Synchronize()
	if err != nil {
		panic(err.Error())
	}
	return nil
}
