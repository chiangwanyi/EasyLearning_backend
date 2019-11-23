package server

import (
	"easy_learning/api"
	"easy_learning/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Sessions())
	r.Use(middleware.Cors())

	r.GET("/ping", api.Ping)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", api.UserRegister)
	}

	return r
}
