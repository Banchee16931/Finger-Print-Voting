package api

import (
	"encoding/json"
	"finger-print-voting-backend/internal/types"
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

			voteMap := make(map[int]int)

			for _, vote := range votes {
				voteMap[vote.CandidateID] = voteMap[vote.CandidateID] + 1
			}

			candidateVotes := []types.CandidateVotes{}
			for _, candidate := range candidates {
				candidateVotes = append(candidateVotes, types.CandidateVotes{
					FirstName:   candidate.FirstName,
					LastName:    candidate.LastName,
					Party:       candidate.Party,
					PartyColour: candidate.PartyColour,
					Votes:       voteMap[candidate.CandidateID],
				})
			}

			electionStates = append(electionStates, types.ElectionState{
				ElectionID: election.ElectionID,
				Start:      election.Start,
				End:        election.End,
				Location:   election.Location,
				Result:     candidateVotes,
			})
		}

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&electionStates)
}
