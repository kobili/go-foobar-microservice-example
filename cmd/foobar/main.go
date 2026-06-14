package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Message struct {
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}

func main() {
	FOO_SERVICE_URL := os.Getenv("FOO_SERVICE_URL")
	BAR_SERVICE_URL := os.Getenv("BAR_SERVICE_URL")

	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"message": "hello foobar",
		})
	})

	mux.HandleFunc("GET /foobar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		fooResponse, err := httpClient.Get(fmt.Sprintf("%sfoo", FOO_SERVICE_URL))
		if err != nil {
			writeError(w, err)
			return
		}
		defer fooResponse.Body.Close()

		fooResponseBody, err := io.ReadAll(fooResponse.Body)
		if err != nil {
			writeError(w, err)
			return
		}

		var fooMessage Message
		err = json.Unmarshal(fooResponseBody, &fooMessage)
		if err != nil {
			writeError(w, err)
			return
		}

		barResponse, err := httpClient.Get(fmt.Sprintf("%sbar", BAR_SERVICE_URL))
		if err != nil {
			writeError(w, err)
			return
		}
		defer barResponse.Body.Close()

		barResponseBody, err := io.ReadAll(barResponse.Body)
		if err != nil {
			writeError(w, err)
			return
		}

		var barMessage Message
		err = json.Unmarshal(barResponseBody, &barMessage)
		if err != nil {
			writeError(w, err)
			return
		}

		message := Message{
			Message: fmt.Sprintf("%s%s", fooMessage.Message, barMessage.Message),
		}

		json.NewEncoder(w).Encode(message)
	})

	SERVER_PORT := os.Getenv("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "8081"
	}

	fmt.Printf("listening on port %s\n", SERVER_PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", SERVER_PORT), mux)
}
