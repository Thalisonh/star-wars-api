package planets

import "github.com/Thalisonh/star-wars-api/models"

type PlanetsServiceSpy struct {
	IPlanetsService
	CreateResponse    *models.Planets
	CreateError       error
	GetAllResponse    *[]models.Planets
	GetAllError       error
	GetByIdResponse   *models.Planets
	GetByIdError      error
	GetByNameResponse *models.Planets
	GetByNameError    error
}

func (p *PlanetsServiceSpy) Create(planet *models.Planets) (*models.Planets, error) {
	return p.CreateResponse, p.CreateError
}

func (p *PlanetsServiceSpy) GetAll() (*[]models.Planets, error) {
	return p.GetAllResponse, p.GetAllError
}

func (p *PlanetsServiceSpy) GetById(planetId int) (*models.Planets, error) {
	return p.GetByIdResponse, p.GetByIdError
}

func (r *PlanetsServiceSpy) GetByName(planetName string) (*models.Planets, error) {
	return r.GetByNameResponse, r.GetByNameError
}
