package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"net/http"
	"time"
)

func (srv *Server) HandleGetElectionForUser(w http.ResponseWriter, r *http.Request) {
	userCtx := r.Context().Value(types.UserContext)

	user, ok := userCtx.(types.UserDetails)
	if !ok {
		HTTPError(w, http.StatusInternalServerError, fmt.Errorf("missing user details"))
		return
	}

	elections, statusCode, err := ElectionsFromUser(srv, user.Username)
	if HTTPError(w, statusCode, err) {
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&elections[0])
}

func ElectionsFromUser(srv *Server, username string) ([]types.Election, int, error) {
	voter, err := srv.db.GetVoter(username)
	if err != nil {
		return []types.Election{}, http.StatusInternalServerError, err
	}

	// get the elections
	elections, err := srv.db.GetElectionByLocation(voter.Location)
	if err != nil {
		return []types.Election{}, http.StatusInternalServerError, err
	}

	now := time.Now()

	// filtering out old elections
	positive := []types.Election{}
	for i := range elections {
		endDate, err := types.StringToDate(elections[i].End)
		if err != nil {
			return []types.Election{}, http.StatusInternalServerError, err
		}

		if endDate.Unix() > now.Unix() {
			positive = append(positive, elections[i])
		}
	}

	if len(positive) <= 0 {
		return []types.Election{}, http.StatusInternalServerError, fmt.Errorf("no elections running")
	}

	return positive, 200, nil
}
