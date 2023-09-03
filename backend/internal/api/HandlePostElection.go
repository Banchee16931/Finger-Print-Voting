package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
	"net/http"
	"time"
)

func (srv *Server) HandlePostElection(w http.ResponseWriter, r *http.Request) {
	var electionReq types.ElectionRequest

	err := json.NewDecoder(r.Body).Decode(&electionReq)
	if HTTPError(w, http.StatusBadRequest, err) {
		return
	}

	if err := electionReq.Validate(); HTTPError(w, http.StatusBadRequest, err) {
		return
	}

	// get the elections
	electionsWithSameLocation, err := srv.db.GetElectionByLocation(electionReq.Location)
	if HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	now := time.Now()

	// filtering out old elections
	existingElections := []types.Election{}
	for i := range electionsWithSameLocation {
		endDate, err := types.StringToDate(electionsWithSameLocation[i].End)
		if HTTPError(w, http.StatusInternalServerError, err) {
			return
		}

		if endDate.Unix() > now.Unix() {
			existingElections = append(existingElections, electionsWithSameLocation[i])
		}
	}

	// checking if any elections with that location are still ongoing
	if len(existingElections) > 0 {
		HTTPError(w, http.StatusBadRequest, fmt.Errorf("election for that location already exists"))
		return
	}

	if err = srv.db.StoreElection(electionReq); HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	log.Println("Stored Election")

	w.WriteHeader(http.StatusCreated)
}
