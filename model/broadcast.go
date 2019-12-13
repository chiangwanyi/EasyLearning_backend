package model

import (
	"easy_learning/db"
	"github.com/globalsign/mgo/bson"
	"time"
)

// Broadcast 留言板模型
type Broadcast struct {
	Id        bson.ObjectId `bson:"_id"`
	FromId    bson.ObjectId `bson:"fromId"`
	Title     string        `bson:"title"`
	Content   string        `bson:"content"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}

func (broadcast *Broadcast) CreateBroadcast() error{
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("broadcast")

	broadcast.Id = bson.NewObjectId()
	broadcast.CreatedAt = time.Now()
	broadcast.UpdatedAt = broadcast.CreatedAt
	return client.Insert(broadcast)
}

func FindBroadcastByFromId(uid string) (list []Broadcast, err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("broadcast")

	if err = client.Find(bson.M{"fromId": bson.ObjectIdHex(uid)}).All(&list); err == nil {
		return list, nil
	} else {
		return []Broadcast{}, err
	}
}

func ShowAllBroadcast() (list []Broadcast, err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("broadcast")

	if err = client.Find(bson.M{}).All(&list); err == nil {
		return list, nil
	} else {
		return []Broadcast{}, err
	}
}
