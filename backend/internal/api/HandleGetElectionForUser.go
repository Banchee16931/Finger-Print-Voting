package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
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
	json.NewEncoder(w).Encode(&elections)
}

func ElectionsFromUser(srv *Server, username string) (types.ElectionWithCandidates, int, error) {
	voter, err := srv.db.GetVoter(username)
	if err != nil {
		return types.ElectionWithCandidates{}, http.StatusInternalServerError, err
	}

	// get the elections
	elections, err := srv.db.GetElectionByLocation(voter.Location)
	if err != nil {
		return types.ElectionWithCandidates{}, http.StatusInternalServerError, err
	}

	now := time.Now()

	// filtering out old elections
	positive := []types.Election{}
	for i := range elections {
		endDate, err := types.StringToDate(elections[i].End)
		if err != nil {
			return types.ElectionWithCandidates{}, http.StatusInternalServerError, err
		}

		if endDate.Unix() > now.Unix() {
			positive = append(positive, elections[i])
		}
	}

	if len(positive) <= 0 {
		return types.ElectionWithCandidates{}, http.StatusInternalServerError, fmt.Errorf("no elections running")
	}

	log.Println("Getting Candidates for: ", positive[0].ElectionID)
	candidates, err := srv.db.GetCandidates(positive[0].ElectionID)
	if err != nil {
		return types.ElectionWithCandidates{}, http.StatusInternalServerError, fmt.Errorf("failed to get candidates")
	}

	if len(candidates) < 2 {
		return types.ElectionWithCandidates{}, http.StatusInternalServerError, fmt.Errorf("no candidates in the election")
	}

	return types.ElectionWithCandidates{
		ElectionID: positive[0].ElectionID,
		Start:      positive[0].Start,
		End:        positive[0].End,
		Location:   positive[0].Location,
		Candidates: candidates,
	}, 200, nil
}
