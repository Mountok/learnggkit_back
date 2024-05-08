package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type UserStorage struct {
	databasePool *pgxpool.Pool
}

func NewUserStorage(databasePool *pgxpool.Pool) *UserStorage {
	storage := new(UserStorage)
	storage.databasePool = databasePool
	return storage
}

func (db *UserStorage) ChangeUserName(userId int, newName string) (userName string,err error) {
	query := "update profiles set full_name = $1 where user_id = $2";
	_, err = db.databasePool.Exec(context.Background(),query,newName,userId)
	logrus.Println("обновление данных в бд")
	if err != nil {
		logrus.Println("ошибка при обновлении")
		return "",err
	}
	logrus.Println("все прошло успешло")
	return newName, nil
}