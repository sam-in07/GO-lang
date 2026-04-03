package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() //2:47:20
		log.Println("ami middlware : age oprint hobo je ")
		next.ServeHTTP(w, r) //10   //middleware kintu
		log.Println("ami middlearw ami pore print hobo")
		diff := time.Since(start)
		log.Println(r.Method, r.URL.Path, diff)

	})
	return handler

}
