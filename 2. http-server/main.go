package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})
	fs := http.FileServer(http.Dir("image/1.jpg"))
	http.Handle("/image/", http.StripPrefix("/image/", fs))
	http.ListenAndServe(":8080", nil)
}
