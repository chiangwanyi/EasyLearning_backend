package model

import (
	"easy_learning/db"
	"github.com/globalsign/mgo/bson"
	"time"
)

// Exam 测试题模型
type Exam struct {
	Id          bson.ObjectId `bson:"_id"`
	ExamName    string        `bson:"examName"`
	Description string        `bson:"description"`
	Deadline    time.Time     `bson:"deadline"`
	Size        int           `bson:"size"`
	Score       float64       `bson:"score"`

	/*
		Questions 测试题的题目
		内容格式：
		{
			"content": ["第一题", "第二题", "第三题"]
		}
	*/
	Questions string `bson:"questions"`

	/*
		Options 测试题的选项
		内容格式：
		{
			"content": [{
				"item": ["选项1的描述", "选项2的描述", "选项3的描述", "选项4的描述"]
			}, {
				"item": ["选项1的描述", "选项2的描述", "选项3的描述", "选项4的描述"]
			}, {
				"item": ["选项1的描述", "选项2的描述", "选项3的描述", "选项4的描述"]
			}]
		}
	*/
	Options string `bson:"options"`

	/*
		Keys 测试题的答案
		内容格式：
		{
			"content": ["A", "B", "C"]
		}
	*/
	Keys string `bson:"keys"`

	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func (exam *Exam) CreateExam() error {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("exam")

	exam.Id = bson.NewObjectId()
	exam.CreatedAt = time.Now()
	exam.UpdatedAt = exam.CreatedAt

	return client.Insert(exam)
}

func FindExamById(id string) (exam Exam, err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("exam")

	if err = client.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&exam); err == nil {
		return exam, nil
	} else {
		return Exam{}, err
	}
}

func ShowAllExam() (examlist []Exam, err error) {
	session := db.MongoSession.Copy()
	defer session.Close()
	client := session.DB("").C("exam")

	if err = client.Find(bson.M{}).All(&examlist); err == nil {
		return examlist, nil
	} else {
		return []Exam{}, err
	}
}
