package main

import (
	"html/template"
	"net/http"
)

type ContactDetails struct {
	Name    string
	Email   string
	Message string
}

func main() {
	temp := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			temp.Execute(writer, nil)
			return
		}
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Name:    r.FormValue("name"),
			Message: r.FormValue("message"),
		}
		_ = details

		temp.Execute(writer, struct{ Success bool }{true})
	})
	http.ListenAndServe(":8080", nil)
}
