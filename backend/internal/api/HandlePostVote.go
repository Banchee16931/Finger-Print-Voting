package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"finger-print-voting-backend/internal/fingerprints"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Votes for a given election
func (srv *Server) HandlePostVote(w http.ResponseWriter, r *http.Request) {
	userCtx := r.Context().Value(types.UserContext) // get user account

	user, ok := userCtx.(types.UserDetails)
	if !ok {
		HTTPError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user account"))
		return
	}

	voter, err := srv.db.GetVoter(user.Username)
	if err != nil {
		HTTPError(w, http.StatusForbidden, fmt.Errorf("user is not a voter"))
		return
	}

	var voteReq types.VoteRequest

	err = json.NewDecoder(r.Body).Decode(&voteReq)
	if HTTPError(w, http.StatusBadRequest, err) {
		return
	}

	if err := voteReq.Validate(); HTTPError(w, http.StatusBadRequest, err) {
		return
	}

	election, err := srv.db.GetElection(voteReq.ElectionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			HTTPError(w, http.StatusForbidden, fmt.Errorf("election doesn't exist"))
			return
		}

		HTTPError(w, http.StatusForbidden, err)
		return
	}
	if HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	if election.Location != voter.Location {
		HTTPError(w, http.StatusBadRequest, fmt.Errorf("user is not from this elections location"))
		return
	}

	startDate, err := types.StringToDate(election.Start)
	if HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	endDate, err := types.StringToDate(election.End)
	if HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	now := time.Now()

	if startDate.Unix() > now.Unix() {
		HTTPError(w, http.StatusBadRequest, fmt.Errorf("election has yet to begin"))
		return
	}

	if endDate.Unix() < now.Unix() {
		HTTPError(w, http.StatusBadRequest, fmt.Errorf("election has already passed"))
		return
	}

	matchingFingerprint, err := fingerprints.CompareBase64Fingerprints(voter.Fingerprint, voteReq.Fingerprint)
	if HTTPError(w, http.StatusBadRequest, err) {
		return
	}

	if !matchingFingerprint {
		HTTPError(w, http.StatusBadRequest, fmt.Errorf("given fingerprint is not registered to this account"))
		return
	}

	tx, err := srv.db.Begin()
	if HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	err = srv.db.StoreVote(tx, types.Vote{
		ElectionID:  voteReq.ElectionID,
		CandidateID: voteReq.CandidateID,
		Username:    user.Username,
	})
	if HTTPErrorWithRollback(tx, w, http.StatusInternalServerError, err) {
		return
	}

	err = srv.db.DeleteVoter(tx, user.Username)
	if HTTPErrorWithRollback(tx, w, http.StatusInternalServerError, err) {
		return
	}

	if HTTPError(w, http.StatusInternalServerError, tx.Commit()) {
		return
	}

	log.Println("Stored Election")

	w.WriteHeader(http.StatusCreated)
}
