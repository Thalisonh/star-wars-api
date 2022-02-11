package planets

import (
	"github.com/Thalisonh/star-wars-api/models"
	"github.com/gin-gonic/gin"
)

type IPlanetsController interface {
	Create(c *gin.Context)
}

type PlanetsController struct {
	IPlanetsService
}

func NewPlanetsController(service IPlanetsService) IPlanetsController {
	return &PlanetsController{service}
}

func (p *PlanetsController) Create(c *gin.Context) {
	planet := &models.Planets{}
	c.ShouldBindJSON(planet)

	p.IPlanetsService.Create(planet)
}
