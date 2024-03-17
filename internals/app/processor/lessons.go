package processor

import (
	"fmt"
	"ggkit_learn_service/internals/app/db"
	"ggkit_learn_service/internals/app/models"
)


type LessonsProcessor struct {
	storage *db.LessonsStorage
}


func NewLessonsProcessor(storage *db.LessonsStorage) *LessonsProcessor {
	processor := new(LessonsProcessor)
	processor.storage = storage
	return processor
}

func (process *LessonsProcessor) GetLessonByIdSubjectAndTheme(subjectId,themeId int) (models.Lesson, error) {
	if subjectId <= 0 {
		return models.Lesson{}, fmt.Errorf("Uncorrect subject id = %d",subjectId)
	} else if themeId <= 0 {
		return models.Lesson{}, fmt.Errorf("Uncorrect theme id = %d",themeId)
	}
	res := process.storage.GetLesson(subjectId,themeId)
	var lessons = models.Lesson{
		Id: res[0].Id,
		Upkeep: `` + res[0].Upkeep,
		ThemeId: res[0].ThemeId,
	}
	fmt.Println(lessons)
	return lessons, nil
}