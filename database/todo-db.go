package database

import (
	"errors"

	"github.com/te6lim/whattodo/models"
)

type TodoDB struct {
	db []models.Todo
}

func (database *TodoDB) InsertTodo(todo models.Todo) {
	database.db = append(database.db, todo)
}

func (database *TodoDB) DeleteTodo(id int) {
	for i := 0; i < len(database.db); i++ {
		if i == id {
			database.db = append(database.db[0:i], database.db[i+1:]...)
			break
		}
	}
}

func (database *TodoDB) UpdateTodo(id int, todo models.Todo) {
	for i := 0; i < len(database.db); i++ {
		if i == id {
			database.db = append(database.db[0:i], database.db[i+1:]...)
			database.db = append(database.db, todo)
			break
		}
	}
}

func (database *TodoDB) GetTodo(id int) (models.Todo, error) {
	for i := 0; i < len(database.db); i++ {
		if i == id {
			return database.db[id], nil
		}
	}
	return models.Todo{}, errors.New("no such object")
}

func (database *TodoDB) GetAllTodos() ([]models.Todo) {
	return database.db
}
