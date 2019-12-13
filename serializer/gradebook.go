package serializer

import (
	"easy_learning/model"
	"time"
)

type Gradebook struct {
	Id         string                 `json:"id"`
	ClassId    string                 `json:"class_id"`
	ExamId     string                 `json:"exam_id"`
	ExamName   string                 `json:"exam_name"`
	StudentId  string                 `json:"student_id"`
	Original   map[string]interface{} `json:"original"`
	FinalScore float64                `json:"final_score"`
	CreatedAt  time.Time              `json:"created_at"`
}

func BuildGradeBook(gradebook model.Gradebook) Gradebook {
	exam, _ := model.FindExamById(gradebook.ExamId.Hex())
	return Gradebook{
		Id:         gradebook.Id.Hex(),
		ClassId:    gradebook.ClassId.Hex(),
		ExamId:     gradebook.ExamId.Hex(),
		ExamName:   exam.ExamName,
		StudentId:  gradebook.StudentId.Hex(),
		Original:   gradebook.Original,
		FinalScore: gradebook.FinalScore,
		CreatedAt:  gradebook.CreatedAt,
	}
}

func BuildGradeBooks(items []model.Gradebook) (gradebookList []Gradebook) {
	for _, val := range items {
		gradebookList = append(gradebookList, BuildGradeBook(val))
	}
	return gradebookList
}
