package handler

import (
	"encoding/json"
	"ggkit_learn_service/internals/app/models"
	"ggkit_learn_service/internals/app/processor"
	"net/http"

	"github.com/gorilla/mux"
)

type LoginHandler struct {
	process *processor.LoginProcessor
}

func NewLoginhandler(processor *processor.LoginProcessor) *LoginHandler {
	handler := new(LoginHandler)
	handler.process = processor
	return handler
}

func (handler *LoginHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		WrapError(w, err)
	}

	err, user_id := handler.process.CreateUser(user)
	if err != nil {
		WrapError(w, err)
	}
	var m = map[string]interface{}{
		"result": "ok",
		"data": user_id,
	}
	WrapOK(w, m)
}

func (handler *LoginHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		WrapError(w, err)
		return
	}

	err, user_id := handler.process.Auth(user)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
		"data": user_id,
	}
	WrapOK(w, m)
}


func (handler *LoginHandler) Profile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["user_id"]

	err, data := handler.process.GetProfileByUserId(userID)
	if err != nil {
		WrapError(w,err)
		return
	}

	m := map[string]interface{}{
		"result": "ok",
		"data": data,
	}

	WrapOK(w,m)
}