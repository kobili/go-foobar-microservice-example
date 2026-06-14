package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"foobar/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"message": "hello foo",
		})
	})

	mux.HandleFunc("GET /foo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"message": "foo",
		})
	})

	SERVER_PORT := os.Getenv("SERVER_PORT")

	fmt.Printf("Listening on port %s\n", SERVER_PORT)
	if SERVER_PORT == "" {
		SERVER_PORT = "8000"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", SERVER_PORT), middleware.RequestLogger(mux))
}
