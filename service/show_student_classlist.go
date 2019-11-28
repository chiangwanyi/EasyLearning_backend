package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
)

type ShowStudentClassList struct {
}

func (service *ShowStudentClassList) ClassList(uid string) serializer.Response {
	var classList []model.Class
	user, _ := model.FindUserById(uid)
	classListId := user.ClassList
	if len(classListId) == 0 {
		return serializer.Response{
			Status: serializer.NotFoundError,
			Data:   nil,
			Msg:    "",
			Error:  "没有班级",
		}
	}
	for _, val := range classListId {
		if class, err := model.FindClassById(val.Hex()); err == nil {
			classList = append(classList, class)
		}
	}
	return serializer.Response{
		Status: serializer.OK,
		Data:   serializer.BuildClasses(classList),
		Msg:    "查询列表成功",
		Error:  "",
	}
}
