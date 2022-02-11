package server

import (
	"github.com/Thalisonh/star-wars-api/api"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()

	api.Router(r)

	return r
}
