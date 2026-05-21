package middleware

import (
	"fmt"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s: %s%s?%s\n", r.Method, r.Host, r.URL.Path, r.URL.RawQuery)
		next.ServeHTTP(w, r)
	})
}
