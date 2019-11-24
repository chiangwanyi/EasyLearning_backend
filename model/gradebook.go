package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// Gradebook 成绩册模型
type Gradebook struct {
	Id        bson.ObjectId `bson:"_id"`
	ClassId   bson.ObjectId `bson:"classId"`
	ExamId    bson.ObjectId `bson:"examId"`
	StudentId bson.ObjectId `bson:"studentId"`

	/*
		Original 学生上传的原始作业答案
		内容格式：
		{
			"content": ["A", "B", "C"]
		}
	*/
	Original string `bson:"original"`

	FinalScore float64   `bson:"finalScore"`
	CreatedAt  time.Time `bson:"createdAt"`
	UpdatedAt  time.Time `bson:"updatedAt"`
}
