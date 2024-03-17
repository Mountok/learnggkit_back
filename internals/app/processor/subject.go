package processor

import (
	"fmt"
	"ggkit_learn_service/internals/app/db"
	"ggkit_learn_service/internals/app/models"
	"strconv"
)

type SubjectProcessor struct {
	storage *db.SubjectStorage
}

func NewSubjectProcessor(storage *db.SubjectStorage) *SubjectProcessor {
	processor := new(SubjectProcessor)
	processor.storage = storage
	return processor
}

func (process *SubjectProcessor) SubjectsList() ([]models.Subject, error) {
	return process.storage.GetAllSubjects(), nil
}
func (process *SubjectProcessor) SubjectById(id string) ([]models.Subject, error) {
	num, err := strconv.Atoi(id)
	if err != nil {
		return []models.Subject{}, fmt.Errorf("uncorrect id - (%d) is not integer",id)
	}
	if num <= 0 {
		return []models.Subject{}, fmt.Errorf("uncorrect id (%d)",id)
	}
	return process.storage.GetSubjectById(num), nil
}
