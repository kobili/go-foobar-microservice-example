package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"foobar/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /bar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"message": "bar",
		})
	})

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", middleware.RequestLogger(mux))
}
