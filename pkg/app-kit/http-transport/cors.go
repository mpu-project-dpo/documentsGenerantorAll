package http_transport

import (
	"github.com/rs/cors"
	"net/http"
)

var initializedCors = cors.New(cors.Options{
	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodHead},
	AllowedHeaders: []string{"*"},
	ExposedHeaders: []string{"*"},
})
