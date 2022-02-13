package planets_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Thalisonh/star-wars-api/api/planets"
	"github.com/Thalisonh/star-wars-api/middlewares"
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

		middleware := &middlewares.SwapiServiceSpy{
			TotalFilmsStatusCode: http.StatusOK,
			TotalFilmsCount:      10,
			TotalFilmsError:      nil,
		}

		w := httptest.NewRecorder()
		bodyBytes, _ := json.Marshal(fakePlanet)

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(bodyBytes))

		planets := planets.NewPlanetsController(&success, middleware)
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
		fail := planets.PlanetsServiceSpy{
			CreateError: gorm.ErrInvalidData,
		}

		middleware := &middlewares.SwapiServiceSpy{
			TotalFilmsStatusCode: http.StatusOK,
			TotalFilmsCount:      10,
			TotalFilmsError:      nil,
		}

		w := httptest.NewRecorder()
		bodyBytes, _ := json.Marshal(fakePlanet)

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(bodyBytes))

		planets := planets.NewPlanetsController(&fail, middleware)
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

	t.Run("Create - Should return 0 when Swapi return error", func(t *testing.T) {
		fail := planets.PlanetsServiceSpy{
			CreateError: gorm.ErrInvalidData,
		}

		middleware := &middlewares.SwapiServiceSpy{
			TotalFilmsStatusCode: http.StatusInternalServerError,
			TotalFilmsCount:      10,
			TotalFilmsError:      errors.New("Unexpected error"),
		}

		w := httptest.NewRecorder()
		bodyBytes, _ := json.Marshal(fakePlanet)

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(bodyBytes))

		planets := planets.NewPlanetsController(&fail, middleware)
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

		planets := planets.NewPlanetsController(&success, nil)
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
		fail := planets.PlanetsServiceSpy{
			GetAllError: gorm.ErrRecordNotFound,
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)

		planets := planets.NewPlanetsController(&fail, nil)
		planets.GetAll(ctx)

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

func TestGetByIdController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	path := "/api/1"

	fakePlanet := &models.Planets{}
	gofakeit.Struct(fakePlanet)

	t.Run("GetById - Should a success", func(t *testing.T) {
		success := planets.PlanetsServiceSpy{
			GetByIdResponse: fakePlanet,
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		planets := planets.NewPlanetsController(&success, nil)
		planets.GetById(ctx)

		body, _ := ioutil.ReadAll(w.Body)

		response := &models.Planets{}
		if err := json.Unmarshal(body, &response); err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("GetById - Should return error when parameter is invalid", func(t *testing.T) {
		fail := planets.PlanetsServiceSpy{
			GetByIdError: errors.New("Id must be a integer"),
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "naboo"})

		planets := planets.NewPlanetsController(&fail, nil)
		planets.GetById(ctx)

		body, _ := ioutil.ReadAll(w.Body)

		response := map[string]string{}
		if err := json.Unmarshal(body, &response); err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, response["message"])
		assert.Contains(t, "Id must be a integer", response["message"])
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("GetById - Should return error when planetId not found", func(t *testing.T) {
		fail := planets.PlanetsServiceSpy{
			GetByIdError: gorm.ErrRecordNotFound,
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		planets := planets.NewPlanetsController(&fail, nil)
		planets.GetById(ctx)

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

func TestGetByNameController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	path := "/api/name/naboo"

	fakePlanet := &models.Planets{}
	gofakeit.Struct(fakePlanet)

	t.Run("GetByName - Should a success", func(t *testing.T) {
		success := planets.PlanetsServiceSpy{
			GetByNameResponse: fakePlanet,
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "name", Value: "naboo"})

		planets := planets.NewPlanetsController(&success, nil)
		planets.GetByName(ctx)

		body, _ := ioutil.ReadAll(w.Body)

		response := &models.Planets{}
		if err := json.Unmarshal(body, &response); err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, response)
		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("GetByName - Should return error when parameter is invalid", func(t *testing.T) {
		fail := planets.PlanetsServiceSpy{
			GetByNameError: errors.New("Id must be a integer"),
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "name", Value: ""})

		planets := planets.NewPlanetsController(&fail, nil)
		planets.GetByName(ctx)

		body, _ := ioutil.ReadAll(w.Body)

		response := map[string]string{}
		if err := json.Unmarshal(body, &response); err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, response["message"])
		assert.Contains(t, gorm.ErrInvalidField.Error(), response["message"])
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

	t.Run("GetByName - Should return error when planetName not found", func(t *testing.T) {
		fail := planets.PlanetsServiceSpy{
			GetByNameError: gorm.ErrRecordNotFound,
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "name", Value: "naboo"})

		planets := planets.NewPlanetsController(&fail, nil)
		planets.GetByName(ctx)

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

func TestDeleteController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	path := "/api/1"

	fakePlanet := &models.Planets{}
	gofakeit.Struct(fakePlanet)

	t.Run("Delete - Should a success", func(t *testing.T) {
		success := planets.PlanetsServiceSpy{
			DeleteError: nil,
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		planets := planets.NewPlanetsController(&success, nil)
		planets.Delete(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})

	t.Run("Delete - Should return error when getById fail", func(t *testing.T) {
		fail := planets.PlanetsServiceSpy{
			GetByIdError: gorm.ErrRecordNotFound,
			DeleteError:  gorm.ErrRecordNotFound,
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		planets := planets.NewPlanetsController(&fail, nil)
		planets.Delete(ctx)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("Delete - Should return error when delete fail", func(t *testing.T) {
		fail := planets.PlanetsServiceSpy{
			DeleteError: gorm.ErrInvalidDB,
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "1"})

		planets := planets.NewPlanetsController(&fail, nil)
		planets.Delete(ctx)

		assert.Equal(t, http.StatusNotFound, w.Result().StatusCode)
	})

	t.Run("Delete - Should return error when parameter is invalid", func(t *testing.T) {
		fail := planets.PlanetsServiceSpy{
			DeleteError: errors.New("Id must be a integer"),
		}

		w := httptest.NewRecorder()

		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = httptest.NewRequest(http.MethodGet, path, nil)
		ctx.Params = append(ctx.Params, gin.Param{Key: "id", Value: "naboo"})

		planets := planets.NewPlanetsController(&fail, nil)
		planets.Delete(ctx)

		body, _ := ioutil.ReadAll(w.Body)

		response := map[string]string{}
		if err := json.Unmarshal(body, &response); err != nil {
			t.Error(err)
			return
		}

		assert.NotNil(t, response["message"])
		assert.Contains(t, "Id must be a integer", response["message"])
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
}
