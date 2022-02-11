package planets

import (
	"github.com/Thalisonh/star-wars-api/repository"
)

type IPlanetsService interface {
	Create()
}

type PlanetsService struct {
	repository.IPlanetsRepository
}

func NewPlanetsService(repository repository.IPlanetsRepository) IPlanetsService {
	return &PlanetsService{repository}
}

func (p *PlanetsService) Create() {

}
