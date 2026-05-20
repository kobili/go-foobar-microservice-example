package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"foobar/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /foo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"message": "foo",
		})
	})

	fmt.Println("Listening on port 8000")
	http.ListenAndServe(":8000", middleware.RequestLogger(mux))
}
