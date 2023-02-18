package usecase

import (
	"testing"

	"github.com/boris-on/hflabs-test-task/internal/service/models"
	"github.com/stretchr/testify/require"
)

func TestUpdateTable(t *testing.T) {
	t.Parallel()

	var codeStatuses = []models.StatusCode{
		{Code: "200", Description: "Запрос успешно обработан"},
		{Code: "400", Description: "Некорректный запрос (невалидный JSON или XML)"},
		{Code: "405", Description: "Запрос сделан с методом, отличным от GET или POST"},
		{Code: "413", Description: "Нарушены ограничения:длина параметра query больше 300 символовили количество ограничений в параметре locations больше 100"},
		{Code: "500", Description: "Произошла внутренняя ошибка сервиса"},
		{Code: "503", Description: "Нет лицензии на запрошенный сервис"},
	}

	t.Run("updating table", func(t *testing.T) {
		var gdocsService, err = NewGdocsService("1x6OwkOgIpkWUiUrdLXQSYfBDUg2nYAGotaWx24z2fB8", "../../../configs/client_secret.json")
		require.NoError(t, err)

		gdocsService.UpdateTable(codeStatuses)
		require.NoError(t, err)
	})
}
