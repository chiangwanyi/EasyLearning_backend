package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// Uploaded 上传的作业数据模型
type Uploaded struct {
	Id         bson.ObjectId `bson:"_id"`
	ClassId    bson.ObjectId `bson:"classId"`
	HomeworkId bson.ObjectId `bson:"homeworkId"`
	StudentId  bson.ObjectId `bson:"studentId"`

	/*
		Original 学生上传的原始作业答案
		内容格式：
		{
			"content": ["A", "B", "C"]
		}
	*/
	Original string `bson:"original"`

	FinalScore int       `bson:"finalScore"`
	CreatedAt  time.Time `bson:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt"`
}
