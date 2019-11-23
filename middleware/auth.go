package middleware

import (
	"easy_learning/config"
	"easy_learning/serializer"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// 检查 UserId 是否存在
		uid := session.Get(config.SessionUserId)
		if uid == nil {
			c.JSON(http.StatusBadRequest, serializer.Response{
				Error: "没有权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
