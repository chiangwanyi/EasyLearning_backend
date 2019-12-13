package api

import (
	"easy_learning/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowExamList(c *gin.Context) {
	var s service.ShowExamList
	res := s.ExamList()
	c.JSON(http.StatusOK, res)
}

func ShowExam(c *gin.Context) {
	var s service.ShowExam
	res := s.Show(c.Param("id"))
	c.JSON(http.StatusOK, res)
}
