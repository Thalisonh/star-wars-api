package planets

import "github.com/gin-gonic/gin"

type IPlanetsController interface {
	Create(c *gin.Context)
}

type PlanetsController struct {
}

func NewPlanetsController() IPlanetsController {
	return &PlanetsController{}
}

func (p *PlanetsController) Create(c *gin.Context) {

}
