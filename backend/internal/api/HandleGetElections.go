package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
	"net/http"
)

func (srv *Server) HandleGetElections(w http.ResponseWriter, r *http.Request) {
	// get the elections
	elections, err := srv.db.GetElections()
	if HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	electionStates := []types.ElectionState{}

	for _, election := range elections {
		results, err := srv.db.GetResults(election.ElectionID)
		if HTTPError(w, http.StatusInternalServerError, err) {
			return
		}

		if len(results) > 0 {
			candidateVotes := []types.CandidateVotes{}

			for _, result := range results {
				candidateVotes = append(candidateVotes, types.CandidateVotes{
					FirstName:   result.FirstName,
					LastName:    result.LastName,
					Party:       result.Party,
					PartyColour: result.PartyColour,
					Votes:       result.Votes,
				})
			}

			electionStates = append(electionStates, types.ElectionState{
				ElectionID: election.ElectionID,
				Start:      election.Start,
				End:        election.End,
				Location:   election.Location,
				Result:     candidateVotes,
			})
		} else {
			candidates, err := srv.db.GetCandidates(election.ElectionID)
			if HTTPError(w, http.StatusInternalServerError, err) {
				return
			}

			votes, err := srv.db.GetVotes(election.ElectionID)
			if HTTPError(w, http.StatusInternalServerError, err) {
				return
			}

			candidateVotes := types.MergeCandidatesAndVotes(candidates, votes)

			electionStates = append(electionStates, types.ElectionState{
				ElectionID: election.ElectionID,
				Start:      election.Start,
				End:        election.End,
				Location:   election.Location,
				Result:     candidateVotes,
			})
		}
	}

	if len(electionStates) <= 0 {
		HTTPError(w, http.StatusNotFound, fmt.Errorf("no elections stored"))
		return
	}

	log.Printf("Returned Elections: %v", electionStates)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&electionStates)
}
