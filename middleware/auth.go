package middleware

import (
	"easy_learning/config"
	"easy_learning/model"
	"easy_learning/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommonAuthRequired 检查普通用户登录权限
func CommonAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// 检查 UserId 是否存在
		uid := session.Get(config.SessionUserId)
		if uid == nil {
			c.JSON(http.StatusBadRequest, serializer.Response{
				Status: serializer.BadAuthError,
				Error:  "没有权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// TeacherAuthRequired 检查老师权限
func TeacherAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		uid := session.Get(config.SessionUserId)
		user, err := model.FindUserById(uid.(string))
		if err != nil || user.Type != "teacher" {
			c.JSON(http.StatusBadRequest, serializer.Response{
				Status: serializer.BadAuthError,
				Error:  "没有老师权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
