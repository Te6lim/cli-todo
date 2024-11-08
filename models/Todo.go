package models

type Todo struct {
	Title     string `json:"title"`
	DateAdded string `json:"dateAdded"`
	Notes     []Note `json:"notes"`
	IsDone    bool   `json:"isDone"`
}

type Note struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	DateAdded string `json:"dateAdded"`
}
