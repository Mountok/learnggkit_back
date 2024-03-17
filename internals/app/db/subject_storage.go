package db

import (
	"context"
	"ggkit_learn_service/internals/app/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type SubjectStorage struct {
	databasePool *pgxpool.Pool
}

func NewSubjectStorage(pool *pgxpool.Pool) *SubjectStorage {
	storage := new(SubjectStorage)
	storage.databasePool = pool
	return storage
}

func (db *SubjectStorage) GetAllSubjects() []models.Subject {
	query := "SELECT id, title, image, description FROM subjects;"

	var result []models.Subject
	err := pgxscan.Select(context.Background(), db.databasePool, &result, query)
	if err != nil {
		log.Errorln(err)
	}
	return result

}
func (db *SubjectStorage) GetSubjectById(id int) []models.Subject {
	query := "SELECT id, title, image, description FROM subjects WHERE id = $1;"

	var result []models.Subject
	err := pgxscan.Select(context.Background(), db.databasePool, &result, query,id)
	if err != nil {
		log.Errorln(err)
	}
	return result

}

