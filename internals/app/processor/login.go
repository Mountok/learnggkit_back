package processor

import (
	"errors"
	"fmt"
	"ggkit_learn_service/internals/app/db"
	"ggkit_learn_service/internals/app/models"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type LoginProcessor struct {
	storage *db.LoginStorage
}

func NewLoginProcessor(storage *db.LoginStorage) *LoginProcessor {
	process := new(LoginProcessor)
	process.storage = storage
	return process
}

func (processor *LoginProcessor) CreateUser(user models.User) (error,int) {
	currentTime := time.Now()

	user.CreateDate = fmt.Sprintf("%.2d.%.2d.%d-%.2d:%.2d", currentTime.Day(), currentTime.Month(), currentTime.Year(), currentTime.Hour(), currentTime.Minute())
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err, 0
	}
	user.Password = string(hashPass)
	err, user_id := processor.storage.CreateUser(user)
	if err != nil {
		return err, 0
	}
	return nil, user_id
}

func (processor *LoginProcessor) Auth(user models.User) (error,int) {
	// vadition
	findUser, err := processor.storage.GetUserByEmail(user)
	if err != nil {
		return errors.New("такого пользователя не сущетвует"),0
	}
	err = bcrypt.CompareHashAndPassword([]byte(findUser[0].Password), []byte(user.Password))
	if err != nil {
		log.Println("Пароль неверный")
		return errors.New("пароль не верный"), 0
	}
	log.Println("Пароль верный")
	log.Println(findUser)
	return nil, findUser[0].Id
}


func (processor *LoginProcessor) GetProfileByUserId(userid string) (error,[]models.Profile) {
	id, err := strconv.Atoi(userid);
	if err != nil {
		return err, []models.Profile{}
	}
	return processor.storage.GetProfileById(id)
}

