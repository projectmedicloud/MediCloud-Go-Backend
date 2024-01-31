package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to MediCloud-Go-Backend!")
	})

	fmt.Println("Server is starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
