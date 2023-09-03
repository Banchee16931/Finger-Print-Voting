package api

import (
	"log"
	"net/http"
)

func (srv *Server) MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Calling: %s", r.URL.String())
		next.ServeHTTP(w, r)
		log.Printf("Leaving: %s", r.URL.String())
	})
}
