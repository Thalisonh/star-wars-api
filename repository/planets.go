package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type IPlanetsRepository interface {
}

type PlanetsRepository struct{ db *mongo.Client }

func NewPlanetsRepository(db *mongo.Client) IPlanetsRepository {
	return &PlanetsRepository{db}
}
