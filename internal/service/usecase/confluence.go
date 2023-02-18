package usecase

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/boris-on/hflabs-test-task/internal/service/models"
)

type ConfluenceService interface {
	GetCodes() ([]models.StatusCode, error)
}

type confluenceService struct {
	Url string
}

func NewConfluenceService(url string) ConfluenceService {
	return &confluenceService{Url: url}
}

func (c confluenceService) GetCodes() ([]models.StatusCode, error) {

	var statusCodes []models.StatusCode

	res, err := http.Get(c.Url)
	if err != nil {
		return []models.StatusCode{}, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		return []models.StatusCode{}, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return []models.StatusCode{}, err
	}

	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		code := s.Find("td").Eq(0).Text()
		description := s.Find("td").Eq(1).Text()

		statusCodes = append(statusCodes, models.StatusCode{
			Code:        code,
			Description: description,
		})
	})

	return statusCodes[1:], nil
}
