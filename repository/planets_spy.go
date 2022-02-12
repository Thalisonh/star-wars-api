package repository

import "github.com/Thalisonh/star-wars-api/models"

type IPlanetsRepositorySpy struct {
	CreateResponse *models.Planets
	CreateError    error
}

func (r *IPlanetsRepositorySpy) Create(planet *models.Planets) (*models.Planets, error) {
	return r.CreateResponse, r.CreateError
}
