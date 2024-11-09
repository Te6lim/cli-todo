package database

import (
	"errors"

	"github.com/te6lim/cli-todo/models"
)

type todoDb struct {
	db []models.Todo
}

var TodoDatabase = todoDb{}

func (database *todoDb) InsertTodo(todo models.Todo) {
	database.db = append(database.db, todo)
}

func (database *todoDb) DeleteTodo(id int64) {
	for i := 0; i < len(database.db); i++ {
		if int64(i) == id {
			database.db = append(database.db[0:i], database.db[i+1:]...)
			break
		}
	}
}

func (database *todoDb) UpdateTodo(id int64, todo models.Todo) {
	for i := 0; i < len(database.db); i++ {
		if int64(i) == id {
			database.db = append(database.db[0:i], database.db[i+1:]...)
			database.db = append(database.db, todo)
			break
		}
	}
}

func (database *todoDb) GetTodo(id int64) (models.Todo, error) {
	for i := 0; i < len(database.db); i++ {
		if int64(i) == id {
			return database.db[id], nil
		}
	}
	return models.Todo{}, errors.New("no such object")
}

func (database *todoDb) GetAllTodos() []models.Todo {
	return database.db
}
