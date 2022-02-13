package planets

import (
	"github.com/Thalisonh/star-wars-api/models"
	"github.com/Thalisonh/star-wars-api/repository"
)

type IPlanetsService interface {
	Create(planet *models.Planets) (*models.Planets, error)
	GetAll() (*[]models.Planets, error)
	GetById(planetId int) (*models.Planets, error)
	GetByName(planetName string) (*models.Planets, error)
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

func (p *PlanetsService) GetAll() (*[]models.Planets, error) {
	planets, err := p.IPlanetsRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return planets, nil
}

func (p *PlanetsService) GetById(planetId int) (*models.Planets, error) {
	planet, err := p.IPlanetsRepository.GetById(planetId)
	if err != nil {
		return nil, err
	}

	return planet, nil
}

func (p *PlanetsService) GetByName(planetName string) (*models.Planets, error) {
	planet, err := p.IPlanetsRepository.GetByName(planetName)
	if err != nil {
		return nil, err
	}

	return planet, nil
}
