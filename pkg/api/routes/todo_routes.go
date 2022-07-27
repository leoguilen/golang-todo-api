package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leoguilen/simple-go-api/pkg/api/handlers"
)

func MapTodoRoutes(r *mux.Router) {
	r.HandleFunc("/v1/todos", handlers.GetAllTodosHandler).Methods(http.MethodGet)
	r.HandleFunc("/v1/todos", handlers.CreateTodoHandler).Methods(http.MethodPost)
	r.HandleFunc("/v1/todos/{todoId}", handlers.GetTodoByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/v1/todos/{todoId}", handlers.UpdateTodoHandler).Methods(http.MethodPut)
	r.HandleFunc("/v1/todos/{todoId}", handlers.DeleteTodoHandler).Methods(http.MethodDelete)
	r.HandleFunc("/v1/todos/{todoId}/new-status", handlers.SetTodoStatusHandler).Methods(http.MethodPatch)
}
