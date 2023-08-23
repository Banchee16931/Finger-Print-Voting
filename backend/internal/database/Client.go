package database

import "finger-print-voting-backend/internal/types"

type Client interface {
	IsSchemaSetup() (bool, error)
	SetupSchema() error

	StoreAdmin(user types.User) error
	StoreVoter(voter types.Voter) error
	GetVoter(username string) (types.Vote, error)
	GetUser(username string) (types.Vote, error)
	DeleteVoter(voter types.Voter) error

	StoreElection(election types.Election) error
	GetElections() ([]types.Election, error)
	DeleteCandidates(electionID int) error

	StoreVote(vote types.Vote) error
	GetVotes(electionID int) ([]types.Vote, error)
	DeleteVotes(electionID int) error

	StoreRegistrant(registrant types.Registrant) error
	GetRegistrants() ([]types.Registrant, error)
	DeleteRegistrant(registrantID int) error

	StoreResult(result types.Result) error
	GetResults(electionID int) ([]types.Result, error)
}
