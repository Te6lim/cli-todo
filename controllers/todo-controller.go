package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/te6lim/cli-todo/controllers/params"
	"github.com/te6lim/cli-todo/database"
	"github.com/te6lim/cli-todo/models"
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
		responseWriter.WriteHeader(http.StatusBadRequest)
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
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, todo); err != nil {
		log.Fatal("unable to convert json to todo object")
	}

	database.TodoDatabase.InsertTodo(*todo)

	res, err := json.Marshal(todo)
	if err != nil {
		log.Fatal("unable to convert todo object ot json")
	}
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(res)
}

func UpdateTodo(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	todoId, err := strconv.ParseInt(mux.Vars(request)[fmt.Sprint(params.TodoId)], 0, 0)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	todo, err := database.TodoDatabase.GetTodo(todoId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		log.Fatal("unable to parse json into todo object")
	}

	database.TodoDatabase.UpdateTodo(todoId, todo)

	res, _ := json.Marshal(todo)
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(res)
}

func DeleteTodo(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	todoId, err := strconv.ParseInt(mux.Vars(request)[fmt.Sprint(params.TodoId)], 0, 0)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	todo, err := database.TodoDatabase.GetTodo(todoId)
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	database.TodoDatabase.DeleteTodo(todoId)

	res, _ := json.Marshal(todo)
	responseWriter.WriteHeader(http.StatusOK)
	responseWriter.Write(res)
}
