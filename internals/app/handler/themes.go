package handler

import (
	"ggkit_learn_service/internals/app/processor"
	"net/http"

	"github.com/gorilla/mux"
)

type ThemesHandler struct {
	processor *processor.ThemesProcessor
}

func NewThemesHandler(processor *processor.ThemesProcessor) *ThemesHandler {
	handler := new(ThemesHandler)
	handler.processor = processor
	return handler
}

func (handler *ThemesHandler) Themes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := handler.processor.ThemesBySubjectId(vars)
	if err != nil {
		WrapError(w, err)
	}
	var m = map[string]interface{}{
		"result": "OK",
		"data":   data,
	}
	WrapOK(w, m)
}
