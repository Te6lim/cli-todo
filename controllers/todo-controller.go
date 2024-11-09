package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/te6lim/whattodo/controllers/params"
	"github.com/te6lim/whattodo/database"
	"github.com/te6lim/whattodo/database/interfaces"
	"github.com/te6lim/whattodo/models"
)

var todoDatabase interfaces.TodoStorage = &database.TodoDb{}

func GetAllTodos(responseWriter http.ResponseWriter, request *http.Request) {

}

func GetTodo(responseWriter http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	todoId, err := strconv.ParseInt(parameters[params.TodoId], 0, 0)
	if err != nil {
		fmt.Println("unable to parse json")
	}
	todo, err := todoDatabase.GetTodo(todoId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	}
	res, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("unable to convert todo object to json")
	}
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(res)
}

func CreateTodo(responseWriter http.ResponseWriter, request *http.Request) {
	todo := &models.Todo{}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("error reading request body")
	}

	if err := json.Unmarshal(body, todo); err != nil {
		fmt.Println("unable to convert json to todo object")
	}

	todoDatabase.InsertTodo(*todo)

	res, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("unable to convert todo object ot json")
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(res)
}

func UpdateTodo(responseWriter http.ResponseWriter, request *http.Request) {

}

func DeleteTodo(responseWriter http.ResponseWriter, request *http.Request) {

}
