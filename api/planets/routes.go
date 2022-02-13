package planets

import (
	"github.com/Thalisonh/star-wars-api/database"
	"github.com/Thalisonh/star-wars-api/repository"
	"github.com/gin-gonic/gin"
)

func Router(allGroup *gin.RouterGroup) {
	r := repository.NewPlanetsRepository(database.GetDb())
	s := NewPlanetsService(r)
	c := NewPlanetsController(s)

	allGroup.POST("/create", c.Create)
	allGroup.GET("/", c.GetAll)
	allGroup.GET("/:id", c.GetById)
	allGroup.GET("/name/:name", c.GetByName)
	allGroup.DELETE("/:id", c.Delete)
}
