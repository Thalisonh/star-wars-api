package planets

import (
	"github.com/Thalisonh/star-wars-api/models"
	"github.com/Thalisonh/star-wars-api/repository"
)

type IPlanetsService interface {
	Create(planet *models.Planets) (*models.Planets, error)
}

type PlanetsService struct {
	repository.IPlanetsRepository
}

func NewPlanetsService(repository repository.IPlanetsRepository) IPlanetsService {
	return &PlanetsService{repository}
}

func (p *PlanetsService) Create(planet *models.Planets) (*models.Planets, error) {
	newPlanet, err := p.IPlanetsRepository.Create(planet)
	if err != nil {
		return nil, err
	}

	return newPlanet, nil
}
