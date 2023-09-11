package main

import (
	"html/template"
	"net/http"
)

type ToDo struct {
	Title string
	Done  bool
}

type ToDoPageData struct {
	PageTitle string
	ToDos     []ToDo
}

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := ToDoPageData{
			PageTitle: "My To Do list",
			ToDos: []ToDo{
				{Title: "Task 1", Done: true},
				{Title: "Task 2", Done: false},
				{Title: "Task 3", Done: true},
				{Title: "Task 4", Done: false},
			},
		}
		tmpl.Execute(writer, data)
	})
	http.ListenAndServe(":8080", nil)
}
