package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
	"log"
	"net/http"
)

func (srv *Server) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var registrationReq types.RegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&registrationReq)
	if err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := registrationReq.Validate(); err != nil {
		WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err = srv.db.StoreRegistrant(registrationReq); err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}

	log.Println("Stored Registrant")

	w.WriteHeader(http.StatusCreated)
}
