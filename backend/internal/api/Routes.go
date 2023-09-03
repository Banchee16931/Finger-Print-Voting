package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (srv *Server) Route() {
	srv.router = mux.NewRouter()

	srv.router.HandleFunc("/login", srv.HandlePostLogin).Methods(http.MethodPost)
	srv.router.HandleFunc("/registrations", srv.HandlePostRegistration).Methods(http.MethodPost)

	authLoggedIn := srv.router.NewRoute().Subrouter()
	authLoggedIn.Use(srv.MiddlewareAuth(AuthLoggedIn))

	authLoggedIn.HandleFunc("/users", srv.HandleGetUser).Methods(http.MethodGet)

	authVoter := srv.router.NewRoute().Subrouter()
	authVoter.Use(srv.MiddlewareAuth(AuthVoter))

	authAdmin := srv.router.NewRoute().Subrouter()
	authAdmin.Use(srv.MiddlewareAuth(AuthAdmin))

}
