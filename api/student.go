package api

import (
	"easy_learning/config"
	"easy_learning/serializer"
	"easy_learning/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ShowStudentClassList 显示用户（学生）加入的班级接口
func ShowStudentClassList(c *gin.Context) {
	session := sessions.Default(c)

	s := service.ShowStudentClassList{}
	res := s.ClassList(session.Get(config.SessionUserId).(string))
	c.JSON(http.StatusOK, res)
}

// StudentJoinClass 加入班级接口
func StudentJoinClass(c *gin.Context) {
	var s service.UserJoinClassService

	if err := c.ShouldBindJSON(&s); err == nil {
		session := sessions.Default(c)
		if err := s.JoinClass(session.Get(config.SessionUserId).(string)); err != nil {
			c.JSON(http.StatusOK, err)
			return
		}
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.OK,
			Data:   nil,
			Msg:    "加入班级成功",
			Error:  "",
		})
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.BadRequestError,
			Data:   err,
			Msg:    "",
			Error:  "输入数据不匹配",
		})
	}
}

func ShowAllStudent(c *gin.Context) {
	var s service.ShowStudentListService

	res := s.ShowStudentList()
	c.JSON(http.StatusOK, res)
}