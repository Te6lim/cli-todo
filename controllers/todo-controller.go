package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/te6lim/whattodo/controllers/params"
	"github.com/te6lim/whattodo/database"
	"github.com/te6lim/whattodo/database/interfaces"
)

var todoDatabase interfaces.TodoStorage = &database.TodoDb{}

func GetAllTodos(responseWriter http.ResponseWriter, request *http.Request) {

}

func GetTodo(responseWriter http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	todoId, err := strconv.ParseInt(parameters[params.TodoId], 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	todo, err := todoDatabase.GetTodo(todoId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
	}
	res, _ := json.Marshal(todo)
	if err != nil {
		fmt.Println("unable to convert json")
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(res)
}

func CreateTodo(responseWriter http.ResponseWriter, request *http.Request) {

}

func UpdateTodo(responseWriter http.ResponseWriter, request *http.Request) {

}

func DeleteTodo(responseWriter http.ResponseWriter, request *http.Request) {

}
