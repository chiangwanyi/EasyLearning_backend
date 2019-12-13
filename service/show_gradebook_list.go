package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
)

type ShowGradeBookList struct {
}

func (service ShowGradeBookList) ShowList(sid string, cid string) serializer.Response {
	if list, err := model.FindGradeBookListByIds(sid, cid); err == nil {
		return serializer.Response{
			Status: serializer.OK,
			Data:   serializer.BuildGradeBooks(list),
			Msg:    "查询成功",
			Error:  "",
		}
	} else {
		return serializer.Response{
			Status: serializer.NotFoundError,
			Data:   err,
			Msg:    "查询失败",
			Error:  "",
		}
	}
}
