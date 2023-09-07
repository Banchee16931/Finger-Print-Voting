package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"log"
	"net/http"
)

// Gets a list of elections
func (srv *Server) HandleGetElections(w http.ResponseWriter, r *http.Request) {
	// get the elections
	elections, err := srv.db.GetElections()
	if HTTPError(w, http.StatusInternalServerError, err) {
		return
	}

	// the list of elections
	electionStates := []types.ElectionState{}

	for _, election := range elections {
		// gets the results of that election if it exists
		results, err := srv.db.GetResults(election.ElectionID)
		if HTTPError(w, http.StatusInternalServerError, err) {
			return
		}

		if len(results) > 0 { // if the elections has some results
			candidateVotes := []types.CandidateVotes{}

			// convert to []types.CandidateVotes
			for _, result := range results {
				candidateVotes = append(candidateVotes, types.CandidateVotes{
					FirstName:   result.FirstName,
					LastName:    result.LastName,
					Party:       result.Party,
					PartyColour: result.PartyColour,
					Votes:       result.Votes,
				})
			}

			// store the election state
			electionStates = append(electionStates, types.ElectionState{
				ElectionID: election.ElectionID,
				Start:      election.Start,
				End:        election.End,
				Location:   election.Location,
				Result:     candidateVotes,
			})
		} else { // if the election hasn't finished and therefore doesn't have results
			candidates, err := srv.db.GetCandidates(election.ElectionID)
			if HTTPError(w, http.StatusInternalServerError, err) {
				return
			}

			votes, err := srv.db.GetVotes(election.ElectionID)
			if HTTPError(w, http.StatusInternalServerError, err) {
				return
			}

			candidateVotes := types.MergeCandidatesAndVotes(candidates, votes)

			// store the election state
			electionStates = append(electionStates, types.ElectionState{
				ElectionID: election.ElectionID,
				Start:      election.Start,
				End:        election.End,
				Location:   election.Location,
				Result:     candidateVotes,
			})
		}
	}

	// if there are no elections report a StatusNotFound http code
	if len(electionStates) <= 0 {
		HTTPError(w, http.StatusNotFound, fmt.Errorf("no elections stored"))
		return
	}

	log.Printf("Returned Elections: %v", electionStates)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&electionStates) // output the elections
}
