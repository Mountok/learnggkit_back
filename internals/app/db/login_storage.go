package db

import (
	"context"
	"errors"
	"ggkit_learn_service/internals/app/models"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type LoginStorage struct {
	databasePool *pgxpool.Pool
}

func NewLoginStorage(databasePool *pgxpool.Pool) *LoginStorage {
	storage := new(LoginStorage)
	storage.databasePool = databasePool
	return storage
}

func (db *LoginStorage) CreateUser(user models.User) error {
	_, err := db.GetUserByEmail(user)
	if err == nil {
		log.Println("Такой пользователь уже есть")
		return errors.New("пользователь с таким email существует")
	}
	query := "INSERT INTO users (email, password, create_date) values ($1,$2,$3);"
	_, err = db.databasePool.Exec(context.Background(), query, user.Email, user.Password, user.CreateDate)
	if err != nil {
		log.Println("Ошибка sql запроса")
		return err
	}
	log.Println("Пользователь создан")
	err = db.CreateProfileForUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (db *LoginStorage) GetUserByEmail(user models.User) (res []models.User, err error) {
	query := "SELECT id, email, password, create_date FROM users WHERE email = $1"
	err = pgxscan.Select(context.Background(), db.databasePool, &res, query, user.Email)
	if err != nil {
		log.Println("Ошибка sql запроса")
		return res, err
	}
	if len(res) == 0 {
		log.Println("Пользователь не найден")
		return res, errors.New("пользователь не найдет")
	}
	log.Println("Пользователь найден")
	return res, nil
}


func (db *LoginStorage) CreateProfileForUser(user models.User) error {
	query := "insert into profiles (user_id,description,phone,full_name, image) values ($1,$2,$3,$4,$5);"
	log.Println("Получение пользователя по почте")
	currentUser,err := db.GetUserByEmail(user)
	if err != nil {
		return err
	}
	log.Println("Создание профиля для пользователя")
	_, err = db.databasePool.Exec(context.Background(),query,currentUser[0].Id,"-","-","-","admin.png")
	if err != nil {
		return err
	}
	log.Println("Профиль создан")
	return nil
}

// фунуция для получения профиля по id пользователя