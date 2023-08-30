package api

import (
	"finger-print-voting-backend/internal/database"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	server         http.Server
	db             database.Database
	router         *mux.Router
	passwordSecret string
	setupRoutes    sync.Once
}

func NewServer() *Server {
	return &Server{}
}

func (srv *Server) Start(addr string) error {
	srv.setupRoutes.Do(srv.Route)

	srv.server = http.Server{
		Addr:              addr,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       30 * time.Second,
		Handler:           srv,
	}

	log.Println("Server Started!")

	return srv.server.ListenAndServe()
}

func (srv *Server) Close() error {
	return srv.server.Close()
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(w, r)
}

func (srv *Server) WithDBClient(db database.Database) *Server {
	srv.db = db
	return srv
}

func (srv *Server) WithPasswordSecret(secret string) *Server {
	srv.passwordSecret = secret
	return srv
}
