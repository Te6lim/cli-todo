package main

import (
	"fmt"

	"github.com/te6lim/whattodo/models"
	"github.com/te6lim/whattodo/storage"
)

func main() {
	fmt.Println("here we go again...")
	var database storage.TodoStorage = &storage.TodoDB{}

	fmt.Println("inserting into todo db...")
	database.InsertTodo(models.Todo{
		Title: "My first todo", DateAdded: "8/11/2024", IsDone: false,
	})
	database.InsertTodo(models.Todo{
		Title: "My second todo", DateAdded: "8/11/2024", IsDone: false,
	})
	database.InsertTodo(models.Todo{
		Title: "My third todo", DateAdded: "8/11/2024", IsDone: false,
	})
	database.InsertTodo(models.Todo{
		Title: "My fourth todo", DateAdded: "8/11/2024", IsDone: false,
	})

	todo, err := database.GetTodo(2)
	if err == nil {
		fmt.Printf("%v was added\n", todo)
	} else {
		fmt.Println(err)
	}

	position := 2
	fmt.Printf("deleting todo at position %v\n", position)
	database.DeleteTodo(position)

	fmt.Printf("all todos: %v\n", database.GetAllTodos())

	position = 1
	fmt.Printf("updating todo at %v\n", position)
	database.UpdateTodo(position, models.Todo{
		Title: fmt.Sprintf("updated todo for %v", position), DateAdded: "8/11/2024", IsDone: false,
	})

	fmt.Printf("all todos: %v\n", database.GetAllTodos())
}
