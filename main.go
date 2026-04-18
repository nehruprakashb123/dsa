package main

import (
	"fmt"
	"net/http"
)

// Add is a simple function we will test in our pipeline
func Add(a int, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! 2 + 3 is %d", Add(2, 3))
	})

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}