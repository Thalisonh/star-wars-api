package planets

import (
	"net/http"
	"strconv"

	"github.com/Thalisonh/star-wars-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IPlanetsController interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	GetByName(c *gin.Context)
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

func (p *PlanetsController) GetAll(c *gin.Context) {
	planets, err := p.IPlanetsService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": gorm.ErrRecordNotFound.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, planets)
}

func (p *PlanetsController) GetById(c *gin.Context) {
	id := c.Param("id")

	planetId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Id must be a integer",
		})
		return
	}

	planets, err := p.IPlanetsService.GetById(planetId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": gorm.ErrRecordNotFound.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, planets)
}

func (p *PlanetsController) GetByName(c *gin.Context) {
	planetName := c.Param("name")

	if planetName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": gorm.ErrInvalidField.Error(),
		})
		return
	}

	planets, err := p.IPlanetsService.GetByName(planetName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": gorm.ErrRecordNotFound.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, planets)
}
