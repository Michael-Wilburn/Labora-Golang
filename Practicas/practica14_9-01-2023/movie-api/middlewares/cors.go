package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

// GetCorsHandler devuelve un middleware CORS configurado
func GetCorsHandler() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080/"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	})

	return c
}
