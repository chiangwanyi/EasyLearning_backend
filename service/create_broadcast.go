package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
	"github.com/globalsign/mgo/bson"
)

type CreateBroadcastService struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func (service CreateBroadcastService) CreateBroadcast(sid string) serializer.Response {
	broadcast := model.Broadcast{
		FromId:  bson.ObjectIdHex(sid),
		Title:   service.Title,
		Content: service.Content,
	}

	if err := broadcast.CreateBroadcast(); err != nil {
		return serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "创建留言失败",
		}
	}
	return serializer.Response{
		Status: serializer.OK,
		Data:   serializer.BuildBroadcast(broadcast),
		Msg:    "",
		Error:  "创建留言成功",
	}
}
