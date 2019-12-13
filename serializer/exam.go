package serializer

import (
	"easy_learning/model"
	"time"
)

type Exam struct {
	Id          string    `json:"_id"`
	ExamName    string    `json:"exam_name"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Finished    bool      `json:"finished"`
	Size        int       `json:"size"`
	Score       float64   `json:"score"`
	Questions   string    `json:"questions"`
	Options     string    `json:"options"`
}

func BuildExam(exam model.Exam) Exam {
	return Exam{
		Id:          exam.Id.Hex(),
		ExamName:    exam.ExamName,
		Description: exam.Description,
		Deadline:    exam.Deadline,
		Size:        exam.Size,
		Score:       exam.Score,
		Questions:   exam.Questions,
		Options:     exam.Options,
	}
}

func BuildExams(items []model.Exam) (examList []Exam) {
	for _, val := range items {
		examList = append(examList, BuildExam(val))
	}
	return examList
}
