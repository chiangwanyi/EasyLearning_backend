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
	if err := model.InsertUserClassList(uid, service.ClassId); err != nil {
		return &serializer.Response{
			Status: serializer.BadRequestError,
			Data:   err,
			Msg:    "",
			Error:  "加入班级失败，请检查班级ID",
		}
	}
	return nil
}
