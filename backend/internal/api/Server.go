package api

type Server struct {
	server *http.Server // Handles the incoming requests
	r      *mux.Router  // Routes an api path to a handler
}
