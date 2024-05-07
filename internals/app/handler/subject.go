package handler

import (
	"errors"
	"ggkit_learn_service/internals/app/processor"
	"net/http"

	"github.com/gorilla/mux"
)

type SubjectHandler struct {
	processor *processor.SubjectProcessor
}

func NewSubjectHandler(processor *processor.SubjectProcessor) *SubjectHandler {
	handler := new(SubjectHandler)
	handler.processor = processor
	return handler
}

func (handler *SubjectHandler) List(w http.ResponseWriter, r *http.Request) {

	list, err := handler.processor.SubjectsList()
	if err != nil {
		WrapError(w, err)
	}
	var m = map[string]interface{}{
		"result": "ok",
		"data":   list,
	}
	WrapOK(w, m)

}

func (handler *SubjectHandler) One(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := handler.processor.SubjectById(vars["id"])
	if err != nil {
		WrapError(w, err)
	}
	var m = map[string]interface{}{
		"result": "OK",
		"data":   data,
	}
	WrapOK(w, m)

}

func (handler *SubjectHandler) Image(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	imageName := queryParams.Get("id")
	if imageName != "" {
		imagePath := "./images/" + imageName
		WrapOKImage(w, imagePath)
	}
	WrapError(w, errors.New("Имя изображения не указано"))
}
