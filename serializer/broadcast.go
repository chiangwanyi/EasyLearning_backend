package serializer

import (
	"easy_learning/model"
	"time"
)

type Broadcast struct {
	Id             string    `json:"id"`
	FromUsername   string    `json:"from_username"`
	FromUserGender string    `json:"from_user_gender"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
}

func BuildBroadcast(broadcast model.Broadcast) Broadcast {
	user, _ := model.FindUserById(broadcast.FromId.Hex())
	return Broadcast{
		Id:             broadcast.Id.Hex(),
		FromUsername:   user.Username,
		FromUserGender: user.Gender,
		Title:          broadcast.Title,
		Content:        broadcast.Content,
		CreatedAt:      broadcast.CreatedAt,
	}
}

func BuildBroadcasts(items []model.Broadcast) (broadcastList []Broadcast) {
	for _, val := range items {
		broadcastList = append(broadcastList, BuildBroadcast(val))
	}
	return broadcastList
}
