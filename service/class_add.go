package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
	"github.com/globalsign/mgo/bson"
)

// ClassAddService 班级添加服务
type ClassAddService struct {
	Classname   string `json:"classname" binding:"required,min=3,max=20"`
	Description string `json:"description" binding:"required,min=2,max=100"`
}

// AddClass 添加班级
func (service *ClassAddService) AddClass(tid string) (model.Class, *serializer.Response) {
	// 检查班级名称是否有重复
	_, err := model.FindClassByClassname(service.Classname)
	if err == nil {
		return model.Class{}, &serializer.Response{
			Status: serializer.BadRequestError,
			Data:   err,
			Msg:    "",
			Error:  "创建班级失败，请检查班级信息",
		}
	}

	class := model.Class{
		TeacherId:   bson.ObjectIdHex(tid),
		Classname:   service.Classname,
		Description: service.Description,
	}

	if err := class.CreateClass(); err != nil {
		return class, &serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "创建班级失败",
		}
	}

	return class, nil
}
