package server

import (
	"easy_learning/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Sessions())
	r.Use(middleware.Cors())

	return r
}
