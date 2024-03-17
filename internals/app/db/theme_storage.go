package db

import (
	"context"
	"ggkit_learn_service/internals/app/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type ThemesStorage struct {
	databasePool *pgxpool.Pool
}

func NewThemesStorage(databasePool *pgxpool.Pool) *ThemesStorage {
	storage := new(ThemesStorage)
	storage.databasePool = databasePool
	return storage
}

func (db *ThemesStorage) GetThemesBySubjectId(id int) (result []models.Theme) {
	query := "SELECT id, title, description, subject_id FROM themes WHERE subject_id = $1"
	err := pgxscan.Select(context.Background(),db.databasePool,&result,query,id)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}
