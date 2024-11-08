package interfaces

import(
	"github.com/te6lim/whattodo/models"
)

type TodoStorage interface {
	InsertTodo(todo models.Todo)
	DeleteTodo(id int)
	UpdateTodo(id int, todo models.Todo)
	GetTodo(id int) (models.Todo, error)
	GetAllTodos() ([]models.Todo)
}