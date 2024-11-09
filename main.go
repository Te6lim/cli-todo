package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/te6lim/whattodo/routes"
)

func main() {
	fmt.Println("here we go again...")
	router := mux.NewRouter()
	routes.RegisterTodoRoutes(router)
	http.Handle("/", router)

	fmt.Println("listening on port:2020")
	if err := http.ListenAndServe("localhost:2020", router); err != nil {
		log.Fatal("failed to connect to port:2020")
	}
}
