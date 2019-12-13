package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
)

// UserJoinClassService 用户加入班级服务
type UserJoinClassService struct {
	ClassId string `json:"class_id" binding:"required"`
}

func (service *UserJoinClassService) JoinClass(uid string) *serializer.Response {
	if class, err := model.FindClassById(service.ClassId); err == nil {
		if err := model.InsertUserClassList(uid, class.Id.Hex()); err != nil {
			return &serializer.Response{
				Status: serializer.InternalServerError,
				Data:   err,
				Msg:    "",
				Error:  "加入班级失败，内部错误",
			}
		}

	} else {
		return &serializer.Response{
			Status: serializer.NotFoundError,
			Data:   err,
			Msg:    "",
			Error:  "未找到该班级",
		}
	}
	return nil
}
