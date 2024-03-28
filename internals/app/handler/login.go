package handler

import (
	"encoding/json"
	"ggkit_learn_service/internals/app/models"
	"ggkit_learn_service/internals/app/processor"
	"net/http"
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

	err = handler.process.CreateUser(user)
	if err != nil {
		WrapError(w, err)
	}
	var m = map[string]interface{}{
		"result": "ok",
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

	err = handler.process.Auth(user)
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{}{
		"result": "OK",
	}
	WrapOK(w, m)
}