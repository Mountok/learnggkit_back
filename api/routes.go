package api

import (
	"ggkit_learn_service/internals/app/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRoute(
	subjectHandler *handler.SubjectHandler,
	themeHandler *handler.ThemesHandler,
	lessonsHandler *handler.LessonsHandler,
	loginHandler *handler.LoginHandler,
) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/subject", subjectHandler.List).Methods(http.MethodGet)
	router.HandleFunc("/api/subject/{id}", subjectHandler.One).Methods(http.MethodGet)

	router.HandleFunc("/api/themes/{subject_id}", themeHandler.Themes).Methods(http.MethodGet)

	router.HandleFunc("/api/lessons/{subject_id}/{theme_id}", lessonsHandler.Lesson).Methods(http.MethodGet)

	router.HandleFunc("/api/reg", loginHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/api/auth", loginHandler.Auth).Methods(http.MethodGet)

	router.HandleFunc("/images", subjectHandler.Image).Methods(http.MethodGet)

	return router
}
