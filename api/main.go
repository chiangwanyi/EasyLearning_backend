package api

import (
	"easy_learning/serializer"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping 服务器运行状态检查接口
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Data: gin.H{
			"msg": "pong",
		},
		Msg:   "服务器运行正常",
		Error: "",
	})
}
