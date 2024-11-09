package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/te6lim/whattodo/controllers"
	"github.com/te6lim/whattodo/controllers/params"
)

func RegisterTodoRoutes(router *mux.Router) {
	router.HandleFunc("/todos", controllers.GetAllTodos).Methods(http.MethodGet)
	router.HandleFunc(fmt.Sprintf("/todo/{%v}", params.TodoId), controllers.GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/todo", controllers.CreateTodo).Methods(http.MethodPost)
	router.HandleFunc(fmt.Sprintf("/todo/{%v}", params.TodoId), controllers.UpdateTodo).Methods(http.MethodPut)
	router.HandleFunc(fmt.Sprintf("/todo/{%v}", params.TodoId), controllers.DeleteTodo).Methods(http.MethodDelete)
}
