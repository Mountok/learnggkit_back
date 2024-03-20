package app

import (
	"net/http"

	"github.com/rs/cors"
)

func CarsSettings() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedOrigins: []string{
			"http://localhost:5173",
		},
		AllowCredentials:   true,
		AllowedHeaders:     []string{},
		OptionsPassthrough: true,
		ExposedHeaders:     []string{},
		Debug:              false,
	})
	return c
}
