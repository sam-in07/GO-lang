package middleware

import (
	"log"
	"net/http"
)

func Hudai(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("ami hhudie middleware")
			next.ServeHTTP(w,r)
		},
	)
}