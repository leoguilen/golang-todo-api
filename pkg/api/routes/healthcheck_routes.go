package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leoguilen/simple-go-api/pkg/api/handlers"
)

func MapHealthcheckRoutes(r *mux.Router) {
	r.HandleFunc("/health", handlers.HealthCheckHandler).Methods(http.MethodGet)
}
