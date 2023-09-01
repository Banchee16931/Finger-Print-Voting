package database

import "finger-print-voting-backend/internal/types"

type Database interface {
	StoreUser(user types.User) error
	StoreVoter(voter types.Voter) error
	GetVoter(username string) (types.Voter, error)
	GetUser(username string) (types.User, error)
	DeleteVoter(voter types.Voter) error

	StoreElection(election types.ElectionRequest) error
	GetElections() ([]types.Election, error)
	GetCandidates() ([]types.Candidate, error)
	DeleteCandidates(electionID int) error

	StoreVote(vote types.Vote) error
	GetVotes(electionID int) ([]types.Vote, error)
	DeleteVotes(electionID int) error

	StoreRegistrant(registrant types.RegistrationRequest) error
	GetRegistrants() ([]types.Registrant, error)
	DeleteRegistrant(registrantID int) error

	StoreResult(result types.ResultRequest) error
	GetResults(electionID int) ([]types.Result, error)
}
