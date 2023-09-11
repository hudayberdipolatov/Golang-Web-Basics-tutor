package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"Lastname"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/decode", func(writer http.ResponseWriter, request *http.Request) {
		var user User
		json.NewDecoder(request.Body).Decode(&user)

		fmt.Fprintf(writer, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
	})

	http.HandleFunc("/encode", func(writer http.ResponseWriter, request *http.Request) {
		polat := User{
			Firstname: "Hudayberdi",
			Lastname:  "Polatow",
			Age:       25,
		}
		json.NewEncoder(writer).Encode(polat)
	})

	http.ListenAndServe(":8080", nil)
}
