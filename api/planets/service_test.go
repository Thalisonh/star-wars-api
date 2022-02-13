package planets_test

import (
	"testing"

	"github.com/Thalisonh/star-wars-api/api/planets"
	"github.com/Thalisonh/star-wars-api/models"
	"github.com/Thalisonh/star-wars-api/repository"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	fakePlanet := &models.Planets{}
	gofakeit.Struct(fakePlanet)

	t.Run("Create - Should return a success response", func(t *testing.T) {
		success := repository.IPlanetsRepositorySpy{
			CreateResponse: fakePlanet,
		}

		planets := planets.NewPlanetsService(&success)

		planet, err := planets.Create(fakePlanet)

		assert.Nil(t, err)
		assert.NotNil(t, planet)
	})

	t.Run("Create - Should return error when fail to create planet", func(t *testing.T) {
		failed := repository.IPlanetsRepositorySpy{
			CreateError: gorm.ErrInvalidData,
		}

		planets := planets.NewPlanetsService(&failed)

		planet, err := planets.Create(&models.Planets{})

		assert.NotNil(t, err)
		assert.Nil(t, planet)
	})
}

func TestGetAll(t *testing.T) {
	fakePlanet := &models.Planets{}
	gofakeit.Struct(fakePlanet)

	t.Run("GetAll - Should return a success response", func(t *testing.T) {
		success := repository.IPlanetsRepositorySpy{
			GetAllResponse: &[]models.Planets{*fakePlanet},
		}

		planets := planets.NewPlanetsService(&success)

		planet, err := planets.GetAll()

		assert.Nil(t, err)
		assert.NotNil(t, planet)
	})

	t.Run("GetAll - Should return error when fail to getAll planets", func(t *testing.T) {
		failed := repository.IPlanetsRepositorySpy{
			GetAllError: gorm.ErrRecordNotFound,
		}

		planets := planets.NewPlanetsService(&failed)

		planet, err := planets.GetAll()

		assert.NotNil(t, err)
		assert.Nil(t, planet)
	})
}

func TestGetById(t *testing.T) {
	fakePlanet := &models.Planets{}
	gofakeit.Struct(fakePlanet)

	planetId := 1

	t.Run("GetById - Should return a success response", func(t *testing.T) {
		success := repository.IPlanetsRepositorySpy{
			GetByIdResponse: fakePlanet,
		}

		planets := planets.NewPlanetsService(&success)

		planet, err := planets.GetById(planetId)

		assert.Nil(t, err)
		assert.NotNil(t, planet)
	})

	t.Run("GetById - Should return error when fail to GetById planets", func(t *testing.T) {
		failed := repository.IPlanetsRepositorySpy{
			GetByIdError: gorm.ErrRecordNotFound,
		}

		planets := planets.NewPlanetsService(&failed)

		planet, err := planets.GetById(planetId)

		assert.NotNil(t, err)
		assert.Nil(t, planet)
	})
}

func TestGetByName(t *testing.T) {
	fakePlanet := &models.Planets{}
	gofakeit.Struct(fakePlanet)

	planetId := "naboo"

	t.Run("GetByName - Should return a success response", func(t *testing.T) {
		success := repository.IPlanetsRepositorySpy{
			GetByNameResponse: fakePlanet,
		}

		planets := planets.NewPlanetsService(&success)

		planet, err := planets.GetByName(planetId)

		assert.Nil(t, err)
		assert.NotNil(t, planet)
	})

	t.Run("GetByName - Should return error when fail to GetByName planets", func(t *testing.T) {
		failed := repository.IPlanetsRepositorySpy{
			GetByNameError: gorm.ErrRecordNotFound,
		}

		planets := planets.NewPlanetsService(&failed)

		planet, err := planets.GetByName(planetId)

		assert.NotNil(t, err)
		assert.Nil(t, planet)
	})
}

func TestDelete(t *testing.T) {
	planetId := 1

	t.Run("Delete - Should return a success response", func(t *testing.T) {
		success := repository.IPlanetsRepositorySpy{
			DeleteError: nil,
		}

		planets := planets.NewPlanetsService(&success)

		err := planets.Delete(planetId)

		assert.Nil(t, err)
	})

	t.Run("Delete - Should return error when fail to Delete planets", func(t *testing.T) {
		failed := repository.IPlanetsRepositorySpy{
			DeleteError: gorm.ErrRecordNotFound,
		}

		planets := planets.NewPlanetsService(&failed)

		err := planets.Delete(planetId)

		assert.NotNil(t, err)
	})

	t.Run("Delete - Should return error when fail to get planet by id", func(t *testing.T) {
		failed := repository.IPlanetsRepositorySpy{
			GetByIdError: gorm.ErrRecordNotFound,
		}

		planets := planets.NewPlanetsService(&failed)

		err := planets.Delete(planetId)

		assert.NotNil(t, err)
	})
}
