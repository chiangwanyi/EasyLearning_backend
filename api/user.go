package api

import (
	"easy_learning/serializer"
	"easy_learning/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var s service.UserRegisterService

	// 进行输入数据校验
	if err := c.ShouldBindJSON(&s); err == nil {
		// 注册
		if user, err := s.Register(); err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.JSON(http.StatusOK, serializer.Response{
				Data:  serializer.BuildUser(user),
				Msg:   "注册成功",
				Error: "",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Data:  err,
			Msg:   "",
			Error: "输入数据不匹配",
		})
	}
}