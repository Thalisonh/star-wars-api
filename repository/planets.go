package repository

import (
	"github.com/Thalisonh/star-wars-api/models"
	"gorm.io/gorm"
)

type IPlanetsRepository interface {
	Create(planet *models.Planets) (*models.Planets, error)
}

type PlanetsRepository struct{ db *gorm.DB }

func NewPlanetsRepository(db *gorm.DB) IPlanetsRepository {
	return &PlanetsRepository{db}
}

func (r *PlanetsRepository) Create(planet *models.Planets) (*models.Planets, error) {
	return planet, r.db.Create(planet).Error
}
