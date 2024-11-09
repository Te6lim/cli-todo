package database

import (
	"errors"

	"github.com/te6lim/whattodo/models"
)

type TodoDb struct {
	db []models.Todo
}

func (database *TodoDb) InsertTodo(todo models.Todo) {
	database.db = append(database.db, todo)
}

func (database *TodoDb) DeleteTodo(id int64) {
	for i := 0; i < len(database.db); i++ {
		if int64(i) == id {
			database.db = append(database.db[0:i], database.db[i+1:]...)
			break
		}
	}
}

func (database *TodoDb) UpdateTodo(id int64, todo models.Todo) {
	for i := 0; i < len(database.db); i++ {
		if int64(i) == id {
			database.db = append(database.db[0:i], database.db[i+1:]...)
			database.db = append(database.db, todo)
			break
		}
	}
}

func (database *TodoDb) GetTodo(id int64) (models.Todo, error) {
	for i := 0; i < len(database.db); i++ {
		if int64(i) == id {
			return database.db[id], nil
		}
	}
	return models.Todo{}, errors.New("no such object")
}

func (database *TodoDb) GetAllTodos() []models.Todo {
	return database.db
}
