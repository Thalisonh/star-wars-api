package planets_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Thalisonh/star-wars-api/api/planets"
	"github.com/Thalisonh/star-wars-api/models"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	path := "/api/create"

	fakePlanet := &models.Planets{}
	gofakeit.Struct(fakePlanet)

	t.Run("Create - Should a success", func(t *testing.T) {
		success := planets.PlanetsServiceSpy{
			CreateResponse: fakePlanet,
		}

		w := httptest.NewRecorder()
		bodyBytes, _ := json.Marshal(fakePlanet)

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(bodyBytes))

		planets := planets.NewPlanetsController(&success)
		planets.Create(ctx)

		body, _ := ioutil.ReadAll(w.Body)

		response := &models.Planets{}
		if err := json.Unmarshal(body, &response); err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("Create - Should return error when any error occurred", func(t *testing.T) {
		success := planets.PlanetsServiceSpy{
			CreateError: gorm.ErrInvalidData,
		}

		w := httptest.NewRecorder()
		bodyBytes, _ := json.Marshal(fakePlanet)

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(bodyBytes))

		planets := planets.NewPlanetsController(&success)
		planets.Create(ctx)

		body, _ := ioutil.ReadAll(w.Body)

		response := map[string]string{}
		if err := json.Unmarshal(body, &response); err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, response["message"])
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
}

func TestGetAllController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	path := "/api/"

	fakePlanet := &models.Planets{}
	gofakeit.Struct(fakePlanet)

	t.Run("GetAll - Should a success", func(t *testing.T) {
		success := planets.PlanetsServiceSpy{
			GetAllResponse: &[]models.Planets{*fakePlanet},
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)

		planets := planets.NewPlanetsController(&success)
		planets.GetAll(ctx)

		body, _ := ioutil.ReadAll(w.Body)

		response := &[]models.Planets{}
		if err := json.Unmarshal(body, &response); err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("GetAll - Should return error when any error occurred", func(t *testing.T) {
		success := planets.PlanetsServiceSpy{
			CreateError: gorm.ErrInvalidData,
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)

		planets := planets.NewPlanetsController(&success)
		planets.Create(ctx)

		body, _ := ioutil.ReadAll(w.Body)

		response := map[string]string{}
		if err := json.Unmarshal(body, &response); err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, response["message"])
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
}
