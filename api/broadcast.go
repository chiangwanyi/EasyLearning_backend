package api

import (
	"easy_learning/config"
	"easy_learning/serializer"
	"easy_learning/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBroadcast(c *gin.Context) {
	var s service.CreateBroadcastService
	if err := c.ShouldBindJSON(&s); err == nil {
		session := sessions.Default(c)
		res := s.CreateBroadcast(session.Get(config.SessionUserId).(string))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.BadRequestError,
			Data:   err,
			Msg:    "",
			Error:  "输入数据格式不匹配",
		})
	}
}

func ShowAllBroadcast(c *gin.Context) {
	var s service.ShowBroadcastListService
	res := s.ShowList()
	c.JSON(http.StatusOK, res)
}
