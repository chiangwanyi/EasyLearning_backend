package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// Broadcast 留言板模型
type Broadcast struct {
	Id        bson.ObjectId `bson:"_id"`
	FromId    bson.ObjectId `bson:"fromId"`
	ToId      bson.ObjectId `bson:"toId"`
	Content   string        `bson:"content"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}
