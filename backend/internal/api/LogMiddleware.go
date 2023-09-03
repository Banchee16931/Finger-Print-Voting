package api

import (
	"log"
	"net/http"
)

func (srv *Server) MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("")
		log.Printf("Calling: [%s] %s", r.Method, r.URL.String())
		next.ServeHTTP(w, r)
		log.Printf("Leaving: [%s] %s", r.Method, r.URL.String())
	})
}
