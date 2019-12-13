package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
)

type ShowExam struct {
}

func (service ShowExam) Show(id string) serializer.Response {
	if exam, err := model.FindExamById(id); err == nil {
		return serializer.Response{
			Status: serializer.OK,
			Data:   serializer.BuildExam(exam),
			Msg:    "查找成功",
			Error:  "",
		}
	} else {
		return serializer.Response{
			Status: serializer.NotFoundError,
			Data:   nil,
			Msg:    "查找失败",
			Error:  "",
		}
	}
}
