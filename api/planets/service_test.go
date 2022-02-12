package planets_test

import (
	"testing"

	"github.com/Thalisonh/star-wars-api/api/planets"
	"github.com/Thalisonh/star-wars-api/models"
	"github.com/Thalisonh/star-wars-api/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	success := repository.IPlanetsRepositorySpy{
		CreateResponse: &models.Planets{},
	}

	failed := repository.IPlanetsRepositorySpy{
		CreateError: gorm.ErrInvalidData,
	}

	t.Run("Create - Should return a success response", func(t *testing.T) {
		planets := planets.NewPlanetsService(&success)

		planet, err := planets.Create(&models.Planets{})

		assert.Nil(t, err)
		assert.NotNil(t, planet)
	})

	t.Run("Create - Should return error when fail to create planet", func(t *testing.T) {
		planets := planets.NewPlanetsService(&failed)

		planet, err := planets.Create(&models.Planets{})

		assert.NotNil(t, err)
		assert.Nil(t, planet)
	})
}
