package main

import (
	"fmt"
	"net/http"

	v1 "github.com/projectmedicloud/MediCloud-Go-Backend/api/v1" // Adjust the import path according to your module path and directory structure
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to MediCloud-Go-Backend!")
	})

	// Register the API endpoint
	http.HandleFunc("/api/v1/test-data", v1.GetTestTableData)

	fmt.Println("Server is starting...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
