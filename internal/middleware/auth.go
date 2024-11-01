package middleware

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request from: %v", r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
