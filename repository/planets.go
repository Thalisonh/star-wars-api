package repository

import (
	"github.com/Thalisonh/star-wars-api/models"
	"gorm.io/gorm"
)

type IPlanetsRepository interface {
	Create(planet *models.Planets) (*models.Planets, error)
	GetAll() (*[]models.Planets, error)
	GetById(planetId int) (*models.Planets, error)
	GetByName(planetName string) (*models.Planets, error)
}

type PlanetsRepository struct{ db *gorm.DB }

func NewPlanetsRepository(db *gorm.DB) IPlanetsRepository {
	return &PlanetsRepository{db}
}

func (r *PlanetsRepository) Create(planet *models.Planets) (*models.Planets, error) {
	return planet, r.db.Create(planet).Error
}

func (r *PlanetsRepository) GetAll() (*[]models.Planets, error) {
	var planets []models.Planets

	return &planets, r.db.Find(planets).Error
}

func (r *PlanetsRepository) GetById(planetId int) (*models.Planets, error) {
	var planets models.Planets

	return &planets, r.db.Where("id = ?").First(planets).Error
}

func (r *PlanetsRepository) GetByName(planetName string) (*models.Planets, error) {
	var planets models.Planets

	return &planets, r.db.Where("name = ?").First(planets).Error
}
