package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
	"log"
	"net/http"
)

// Stores a user's registration
func (srv *Server) HandlePostRegistration(w http.ResponseWriter, r *http.Request) {
	var registrationReq types.RegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&registrationReq)
	if HTTPError(w, http.StatusBadRequest, err) {
		return
	}

	if err := registrationReq.Validate(); HTTPError(w, http.StatusBadRequest, err) {
		return
	}

	if err = srv.db.StoreRegistrant(registrationReq); HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	log.Println("Stored Registrant")

	w.WriteHeader(http.StatusCreated)
}
