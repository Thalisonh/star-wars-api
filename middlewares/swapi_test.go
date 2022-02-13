package middlewares_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/Thalisonh/star-wars-api/middlewares"
	"github.com/Thalisonh/star-wars-api/models"
	"github.com/Thalisonh/star-wars-api/util"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func TestTotalFilms(t *testing.T) {
	client := &util.HttpClientMock{}

	fakeSwapi := &models.Swapi{}
	gofakeit.Struct(fakeSwapi)

	planetName := "Naboo"

	t.Run("TotalFilms - Should a success", func(t *testing.T) {
		responseBytes, _ := json.Marshal(fakeSwapi)
		responseIoReader := io.NopCloser(bytes.NewBuffer(responseBytes))

		client.StatusCode = http.StatusOK
		client.Body = responseIoReader
		client.Err = nil

		middlewares := middlewares.NewSwapiService(client)
		statusCode, _, err := middlewares.TotalFilms(planetName)

		assert.Equal(t, statusCode, http.StatusOK)
		assert.Nil(t, err)
	})

	t.Run("TotalFilms - Should return error when fail to call API", func(t *testing.T) {
		responseBytes, _ := json.Marshal(fakeSwapi)
		responseIoReader := io.NopCloser(bytes.NewBuffer(responseBytes))

		client.StatusCode = http.StatusInternalServerError
		client.Body = responseIoReader
		client.Err = errors.New("Unexpected error")

		middlewares := middlewares.NewSwapiService(client)
		statusCode, _, err := middlewares.TotalFilms(planetName)

		assert.Equal(t, statusCode, http.StatusInternalServerError)
		assert.NotNil(t, err)
	})

	t.Run("TotalFilms - Should return total == 1", func(t *testing.T) {
		fakeSwapi := &models.Swapi{
			Count: 1,
			Results: []models.Results{
				{Name: "Naboo", Films: []string{"urlNoboo"}},
			},
		}
		responseBytes, _ := json.Marshal(fakeSwapi)
		responseIoReader := io.NopCloser(bytes.NewBuffer(responseBytes))

		client.StatusCode = http.StatusOK
		client.Body = responseIoReader
		client.Err = nil

		middlewares := middlewares.NewSwapiService(client)
		statusCode, total, err := middlewares.TotalFilms(planetName)

		assert.Equal(t, statusCode, http.StatusOK)
		assert.Nil(t, err)
		assert.Equal(t, total, 1)
	})

	t.Run("TotalFilms - Should return error to unmarshal body", func(t *testing.T) {
		fakeSwapi := map[string]string{
			"Count": "1",
		}
		responseBytes, _ := json.Marshal(fakeSwapi)
		responseIoReader := io.NopCloser(bytes.NewBuffer(responseBytes))

		client.StatusCode = http.StatusInternalServerError
		client.Body = responseIoReader
		client.Err = nil

		middlewares := middlewares.NewSwapiService(client)
		statusCode, total, err := middlewares.TotalFilms(planetName)

		assert.Equal(t, statusCode, http.StatusInternalServerError)
		assert.NotNil(t, err)
		assert.Equal(t, total, 0)
	})
}
