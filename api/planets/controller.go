package planets

import (
	"net/http"

	"github.com/Thalisonh/star-wars-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	planets, err := p.IPlanetsService.Create(planet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": gorm.ErrInvalidDB.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, planets)
}
