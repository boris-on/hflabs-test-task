package usecase

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCodes(t *testing.T) {

	t.Parallel()

	var confluenceService = NewConfluenceService("https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999")

	t.Run("Parsing codes", func(t *testing.T) {
		_, err := confluenceService.GetCodes()
		require.NoError(t, err)
	})
}
