package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (srv *Server) Route() {
	srv.router = mux.NewRouter()
	srv.router.Use(srv.MiddlewareLog)

	srv.router.HandleFunc("/login", srv.HandlePostLogin).Methods(http.MethodPost)
	srv.router.HandleFunc("/registrations", srv.HandlePostRegistration).Methods(http.MethodPost)
	srv.router.HandleFunc("/elections", srv.HandleGetElections).Methods(http.MethodGet)

	authLoggedIn := srv.router.NewRoute().Subrouter()
	authLoggedIn.Use(srv.MiddlewareAuth(AuthLoggedIn))

	authLoggedIn.HandleFunc("/users", srv.HandleGetUser).Methods(http.MethodGet)

	authVoter := srv.router.NewRoute().Subrouter()
	authVoter.Use(srv.MiddlewareAuth(AuthVoter))

	authVoter.HandleFunc("/elections/user", srv.HandleGetElectionForUser).Methods(http.MethodGet)
	authVoter.HandleFunc("/vote", srv.HandlePostVote).Methods(http.MethodPost)

	authAdmin := srv.router.NewRoute().Subrouter()
	authAdmin.Use(srv.MiddlewareAuth(AuthAdmin))

	authAdmin.HandleFunc("/registrations/acceptance", srv.HandlePostRegistrationAcceptance).Methods(http.MethodPost)
	authAdmin.HandleFunc("/registrations", srv.HandleGetRegistrations).Methods(http.MethodGet)
	authAdmin.HandleFunc("/elections", srv.HandlePostElection).Methods(http.MethodPost)

}
