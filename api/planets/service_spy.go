package planets

import "github.com/Thalisonh/star-wars-api/models"

type PlanetsServiceSpy struct {
	CreateResponse *models.Planets
	CreateError    error
}

func (p *PlanetsServiceSpy) Create(planet *models.Planets) (*models.Planets, error) {
	return p.CreateResponse, p.CreateError
}
