package server

import (
	"easy_learning/api"
	"easy_learning/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Session 中间件
	r.Use(middleware.Sessions())
	// Cors 跨域访问中间件
	r.Use(middleware.Cors())

	r.GET("/ping", api.Ping)

	user := r.Group("/api/user")
	{

		user.POST("register", api.UserRegister)
		user.POST("login", api.UserLogin)

		auth := user.Use(middleware.CommonAuthRequired())
		{
			auth.DELETE("logout", api.UserLogout)
		}
	}

	class := r.Group("/api/class")
	{
		auth := class.Use(middleware.CommonAuthRequired(), middleware.TeacherAuthRequired())
		{
			auth.POST("add", api.ClassAdd)
		}
	}
	return r
}
