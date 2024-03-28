package processor

import (
	"fmt"
	"ggkit_learn_service/internals/app/db"
	"ggkit_learn_service/internals/app/models"
	"strconv"
)

type ThemesProcessor struct {
	storage *db.ThemesStorage
}

func NewThemesProcessor(storage *db.ThemesStorage) *ThemesProcessor {
	processor := new(ThemesProcessor)
	processor.storage = storage
	return processor
}

func (process *ThemesProcessor) ThemesBySubjectId(req_vars map[string]string) ([]models.Theme, error) {
	num, err := strconv.Atoi(req_vars["subject_id"])
	if err != nil {
		return []models.Theme{}, fmt.Errorf("error: %s", err.Error())
	}
	return process.storage.GetThemesBySubjectId(num), nil

}
