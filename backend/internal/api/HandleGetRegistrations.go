package api

import (
	"encoding/json"
	"net/http"
)

// Gets a list of registrantions
func (srv *Server) HandleGetRegistrations(w http.ResponseWriter, r *http.Request) {
	// get the registrants
	registrants, err := srv.db.GetRegistrants()
	if HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&registrants) // output the registrants
}
