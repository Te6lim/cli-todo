package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/te6lim/whattodo/controllers/params"
	"github.com/te6lim/whattodo/database"
	"github.com/te6lim/whattodo/models"
)

func GetAllTodos(responseWriter http.ResponseWriter, request *http.Request) {
	todos := database.TodoDatabase.GetAllTodos()
	res, _ := json.Marshal(todos)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(res)
}

func GetTodo(responseWriter http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	todoId, err := strconv.ParseInt(parameters[params.TodoId], 0, 0)
	if err != nil {
		log.Fatal("unable to parse json")
		return
	}
	todo, err := database.TodoDatabase.GetTodo(todoId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(todo)
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(res)
}

func CreateTodo(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	todo := &models.Todo{}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal("error reading request body")
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, todo); err != nil {
		log.Fatal("unable to convert json to todo object")
	}

	database.TodoDatabase.InsertTodo(*todo)

	res, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("unable to convert todo object ot json")
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(res)
}

func UpdateTodo(responseWriter http.ResponseWriter, request *http.Request) {

}

func DeleteTodo(responseWriter http.ResponseWriter, request *http.Request) {

}
