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
