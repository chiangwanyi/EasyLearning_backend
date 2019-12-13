package model

import (
	"easy_learning/db"
	"github.com/globalsign/mgo/bson"
	"time"
)

// Gradebook 成绩册模型
type Gradebook struct {
	Id         bson.ObjectId          `bson:"_id"`
	ClassId    bson.ObjectId          `bson:"classId"`
	ExamId     bson.ObjectId          `bson:"examId"`
	StudentId  bson.ObjectId          `bson:"studentId"`
	Original   map[string]interface{} `bson:"original"`
	FinalScore float64                `bson:"finalScore"`
	CreatedAt  time.Time              `bson:"createdAt"`
	UpdatedAt  time.Time              `bson:"updatedAt"`
}

func (gradebook *Gradebook) CreateGradeBook() error {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("gradebook")

	gradebook.Id = bson.NewObjectId()
	gradebook.CreatedAt = time.Now()
	gradebook.UpdatedAt = gradebook.CreatedAt

	return client.Insert(gradebook)
}

func FindGradeBookListByIds(sid string, cid string) (list []Gradebook, err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("gradebook")

	if err = client.Find(bson.M{"$and": []bson.M{bson.M{"studentId": bson.ObjectIdHex(sid)}, bson.M{"classId": bson.ObjectIdHex(cid)}}}).All(&list); err == nil {
		return list, nil
	} else {
		return []Gradebook{}, err
	}
}

