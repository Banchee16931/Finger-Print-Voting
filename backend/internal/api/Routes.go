package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (srv *Server) Route() {
	srv.router = mux.NewRouter()

	srv.router.HandleFunc("/login", srv.HandleLogin).Methods(http.MethodPost)
	srv.router.HandleFunc("/register", srv.HandleRegister).Methods(http.MethodPost)

	auth := srv.router.NewRoute().Subrouter()
	auth.Use(srv.MiddlewareAuth)
	auth.HandleFunc("/user", srv.HandleUser).Methods(http.MethodGet)
}
