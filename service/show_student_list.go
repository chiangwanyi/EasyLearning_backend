package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
)

type ShowStudentListService struct {

}

func (service ShowStudentListService) ShowStudentList() serializer.Response {
	if list, err := model.ShowAllStudent();err == nil {
		return serializer.Response{
			Status: serializer.OK,
			Data:   serializer.BuildStudents(list),
			Msg:    "获取数据成功",
			Error:  "",
		}
	} else {
		return serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "未知错误",
		}
	}
}
