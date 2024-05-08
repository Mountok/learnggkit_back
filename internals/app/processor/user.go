package processor

import (
	"ggkit_learn_service/internals/app/db"
	"strconv"
)

type UserProcessor struct {
	storage *db.UserStorage
}

func NewUserProcessor(storage *db.UserStorage) *UserProcessor {
	process := new(UserProcessor)
	process.storage = storage
	return process
}


func (processor *UserProcessor) ChangeUserName(userId,newName string) (string,error){
	user_id, err := strconv.Atoi(userId);

	if err !=  nil {
		return "",err
	}
	
	return processor.storage.ChangeUserName(user_id,newName)
}