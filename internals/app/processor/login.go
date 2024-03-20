package processor

import (
	"fmt"
	"ggkit_learn_service/internals/app/db"
	"ggkit_learn_service/internals/app/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginProcessor struct {
	storage *db.LoginStorage
}

func NewLoginProcessor(storage *db.LoginStorage) *LoginProcessor {
	process := new(LoginProcessor)
	process.storage = storage
	return process
}

func (processor *LoginProcessor) CreateUser(user models.User) error {
	currentTime := time.Now()

	user.CreateDate = fmt.Sprintf("%.2d.%.2d.%d-%.2d:%.2d", currentTime.Day(), currentTime.Month(), currentTime.Year(), currentTime.Hour(), currentTime.Minute())
	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashPass)
	err = processor.storage.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (processor *LoginProcessor) Auth(user models.User) error {
	// vadition
	findUser, err := processor.storage.GetUserByEmail(user)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(findUser[0].Password), []byte(user.Password))
	if err != nil {
		log.Println("Пароль неверный")
		return err
	}
	log.Println("Пароль верный")
	log.Println(findUser)
	return nil
}