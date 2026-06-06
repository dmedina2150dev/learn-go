package server

import (
	"fmt"
	"net/http"
)

func InitServer() {
	fmt.Println("Server running")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, you've request %s\n", r.URL.Path)
		fmt.Fprintf(w, "Welcome to website")
	})

	http.ListenAndServe(":3000", nil)
}
