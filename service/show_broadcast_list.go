package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
)

type ShowBroadcastListService struct {
}

func (service ShowBroadcastListService) ShowList() serializer.Response {
	if list, err := model.ShowAllBroadcast(); err == nil {
		return serializer.Response{
			Status: serializer.OK,
			Data:   serializer.BuildBroadcasts(list),
			Msg:    "查询所有数据成功",
			Error:  "",
		}
	} else {
		return serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "查询失败",
		}
	}
}

func (service ShowBroadcastListService) ShowListByFromId(sid string) serializer.Response {
	if list, err := model.FindBroadcastByFromId(sid); err == nil {
		return serializer.Response{
			Status: serializer.OK,
			Data:   serializer.BuildBroadcasts(list),
			Msg:    "查询所有数据成功",
			Error:  "",
		}
	} else {
		return serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "查询失败",
		}
	}
}
