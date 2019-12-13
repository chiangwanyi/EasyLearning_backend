package api

import (
	"easy_learning/serializer"
	"easy_learning/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateExam(c *gin.Context) {
	var s service.ExamAddServices

	if err := c.ShouldBindJSON(&s); err == nil {
		if exam, err := s.AddExam(); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, serializer.Response{
				Status: serializer.OK,
				Data:   serializer.BuildExam(exam),
				Msg:    "注册成功",
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
