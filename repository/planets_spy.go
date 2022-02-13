package repository

import "github.com/Thalisonh/star-wars-api/models"

type IPlanetsRepositorySpy struct {
	IPlanetsRepository
	CreateResponse    *models.Planets
	CreateError       error
	GetAllResponse    *[]models.Planets
	GetAllError       error
	GetByIdResponse   *models.Planets
	GetByIdError      error
	GetByNameResponse *models.Planets
	GetByNameError    error
}

func (r *IPlanetsRepositorySpy) Create(planet *models.Planets) (*models.Planets, error) {
	return r.CreateResponse, r.CreateError
}

func (r *IPlanetsRepositorySpy) GetAll() (*[]models.Planets, error) {
	return r.GetAllResponse, r.GetAllError
}

func (r *IPlanetsRepositorySpy) GetById(planetId int) (*models.Planets, error) {
	return r.GetByIdResponse, r.GetByIdError
}

func (r *IPlanetsRepositorySpy) GetByName(planetName string) (*models.Planets, error) {
	return r.GetByNameResponse, r.GetByNameError
}
