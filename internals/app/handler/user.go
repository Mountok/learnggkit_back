package handler

import (
	"ggkit_learn_service/internals/app/processor"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	processor *processor.UserProcessor
}

func NewUserHandler(processor *processor.UserProcessor) *UserHandler {
	handler := new(UserHandler)
	handler.processor = processor
	return handler
}


func (handler *UserHandler) ChangeName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var (
		userIdString = vars["user_id"]
		newName = vars["new_name"]
	)

	newName, err := handler.processor.ChangeUserName(userIdString,newName)
	if err != nil {
		WrapError(w,err)
		return
	}

	var m = map[string]interface{}{
		"result": http.StatusOK,
		"data": newName,
	}
	WrapOK(w,m)
}