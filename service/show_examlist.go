package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
)

type ShowExamList struct {
}

func (service *ShowExamList) ExamList() serializer.Response {
	if examList, err := model.ShowAllExam(); err == nil {
		return serializer.Response{
			Status: serializer.OK,
			Data:   serializer.BuildExams(examList),
			Msg:    "显示所有数据成功",
			Error:  "",
		}
	} else {
		return serializer.Response{
			Status: serializer.NotFoundError,
			Data:   nil,
			Msg:    "没有数据",
			Error:  "",
		}
	}
}
