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

	// 用户相关路由群组
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", api.UserRegister)
		v1.POST("/login", api.UserLogin)

		// 需要登录权限路由群组
		auth := v1.Group("")
		// 需要普通用户登录权限
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("/logout", api.UserLogout)
		}
	}

	return r
}
