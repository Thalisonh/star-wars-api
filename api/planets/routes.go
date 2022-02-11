package planets

import (
	"github.com/gin-gonic/gin"
)

func Router(allGroup *gin.RouterGroup) {
	r := NewPlanetsController()

	allGroup.POST("", r.Create)
}
