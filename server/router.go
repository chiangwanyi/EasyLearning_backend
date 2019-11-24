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

	v1 := r.Group("/api/v1")
	{
		v1.GET("ping", api.Ping)

		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		auth := v1.Use(middleware.CommonAuthRequired())
		{
			auth.DELETE("user/logout", api.UserLogout)
		}

		auth = v1.Use(middleware.CommonAuthRequired(), middleware.TeacherAuthRequired())
		{
			v1.POST("class/add", api.ClassAdd)
		}

	}
	return r
}
