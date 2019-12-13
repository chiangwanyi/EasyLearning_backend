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

	// 基础用户操作
	user := r.Group("/api/user")
	{

		// 用户注册
		user.POST("register", api.UserRegister)

		// 用户登录
		user.POST("login", api.UserLogin)

		// 需要登录权限
		auth := user.Use(middleware.CommonAuthRequired())
		{
			// 用户主页
			auth.GET("home", api.UserHome)

			// 用户登出
			auth.DELETE("logout", api.UserLogout)

			auth.POST("setCurrentClass", api.UserSetCurrentClass)

			// 需要学生权限
			student := auth.Use(middleware.StudentAuthRequired())
			{
				// 显示学生加入的班级
				student.GET("showClass", api.ShowStudentClassList)

				// 学生加入班级
				student.PUT("joinClass", api.StudentJoinClass)

				student.GET("students", api.ShowAllStudent)
			}
		}
	}

	// 班级操作
	class := r.Group("/api/class")
	{
		// 需要登录权限 和 老师权限
		auth := class.Use(middleware.CommonAuthRequired(), middleware.TeacherAuthRequired())
		{
			// 老师添加班级
			auth.POST("create", api.CreateClass)
		}
	}

	exam := r.Group("/api")
	{

		exam.GET("exams/list", api.ShowExamList)

		exam.GET("exam/:id", api.ShowExam)

		auth := exam.Use(middleware.CommonAuthRequired(), middleware.TeacherAuthRequired())
		{
			auth.POST("exams/create", api.CreateExam)
		}
	}

	grade := r.Group("/api")
	{
		auth := grade.Use(middleware.CommonAuthRequired())
		{
			auth.POST("grade/calc", api.CalculateGrade)
			auth.GET("grades", api.ShowGradeList)
		}
	}

	broadcast := r.Group("/api")
	{
		auth := broadcast.Use(middleware.CommonAuthRequired())
		{
			auth.POST("broadcast", api.CreateBroadcast)
			auth.GET("broadcasts", api.ShowAllBroadcast)
		}
	}
	return r
}
