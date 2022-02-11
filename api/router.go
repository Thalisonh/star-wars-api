package api

import (
	"github.com/Thalisonh/star-wars-api/api/planets"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	allGroups := r.Group("/api")

	planets.Router(allGroups)

}
