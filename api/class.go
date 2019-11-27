package api

import (
	"easy_learning/config"
	"easy_learning/serializer"
	"easy_learning/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateClass 创建接口
func CreateClass(c *gin.Context) {
	var s service.ClassAddService

	if err := c.ShouldBindJSON(&s); err == nil {
		session := sessions.Default(c)
		if class, err := s.AddClass(session.Get(config.SessionUserId).(string)); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, serializer.Response{
				Status: serializer.OK,
				Data:   serializer.BuildClass(class),
				Msg:    "创建班级成功",
				Error:  "",
			})
		}
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.BadRequestError,
			Data:   err,
			Msg:    "",
			Error:  "输入数据不匹配",
		})
	}
}
