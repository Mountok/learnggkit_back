package db

import (
	"context"
	"ggkit_learn_service/internals/app/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)


type LessonsStorage struct {
	databasePool *pgxpool.Pool
}

func NewLessonsStorage(pool *pgxpool.Pool) *LessonsStorage{
	storage := new(LessonsStorage);
	storage.databasePool = pool
	return storage
}

func (db *LessonsStorage) GetLesson(subjectId, themeId int) []models.Lesson {
	var result []models.Lesson
	query := "select * from lessons where theme_id in (select id from themes where subject_id = $1 and id = $2);"
	err := pgxscan.Select(context.Background(),db.databasePool,&result,query,subjectId,themeId)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}