package repository

import "github.com/Thalisonh/star-wars-api/models"

type IPlanetsRepositorySpy struct {
	IPlanetsRepository
	CreateResponse *models.Planets
	CreateError    error
	GetAllResponse *[]models.Planets
	GetAllError    error
}

func (r *IPlanetsRepositorySpy) Create(planet *models.Planets) (*models.Planets, error) {
	return r.CreateResponse, r.CreateError
}

func (r *IPlanetsRepositorySpy) GetAll() (*[]models.Planets, error) {
	return r.GetAllResponse, r.GetAllError
}
