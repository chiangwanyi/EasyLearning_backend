package api

import (
	"easy_learning/config"
	"easy_learning/serializer"
	"easy_learning/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CalculateGrade(c *gin.Context) {
	var s service.CalcGrade
	if err := c.ShouldBindJSON(&s); err == nil {
		session := sessions.Default(c)
		res := s.CalcExamGrade(session.Get(config.SessionUserId).(string), session.Get(config.SessionClassId).(string))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.BadRequestError,
			Data:   err,
			Msg:    "",
			Error:  "输入数据不匹配",
		})
	}
}

func ShowGradeList(c *gin.Context) {
	session := sessions.Default(c)
	var s service.ShowGradeBookList
	res := s.ShowList(session.Get(config.SessionUserId).(string), session.Get(config.SessionClassId).(string))
	c.JSON(http.StatusOK, res)
}
