package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/te6lim/whattodo/controllers"
)

func RegisterTodoRoutes(router *mux.Router) {
	router.HandleFunc("/todos", controllers.GetAllTodos).Methods(http.MethodGet)
	router.HandleFunc("/todo/{todoId}", controllers.GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/todo", controllers.CreateTodo).Methods(http.MethodPost)
	router.HandleFunc("/todo/{todoId}", controllers.UpdateTodo).Methods(http.MethodPut)
	router.HandleFunc("/todo/{todoId}", controllers.DeleteTodo).Methods(http.MethodDelete)
}