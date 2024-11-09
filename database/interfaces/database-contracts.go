package interfaces

import(
	"github.com/te6lim/cli-todo/models"
)

type TodoStorage interface {
	InsertTodo(todo models.Todo)
	DeleteTodo(id int64)
	UpdateTodo(id int64, todo models.Todo)
	GetTodo(id int64) (models.Todo, error)
	GetAllTodos() ([]models.Todo)
}