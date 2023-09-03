package database

import (
	"database/sql"
	"finger-print-voting-backend/internal/types"
)

type Database interface {
	StoreUser(tx *sql.Tx, user types.User) error
	StoreVoter(tx *sql.Tx, voter types.Voter) error
	GetVoter(username string) (types.Voter, error)
	GetUser(username string) (types.User, error)
	DeleteVoter(voter string) error

	StoreElection(election types.ElectionRequest) error
	GetElections() ([]types.Election, error)
	GetCandidates(electionID int) ([]types.Candidate, error)
	DeleteCandidates(electionID int) error

	StoreVote(vote types.Vote) error
	GetVotes(electionID int) ([]types.Vote, error)
	DeleteVotes(electionID int) error

	StoreRegistrant(registrant types.RegistrationRequest) error
	GetRegistrants() ([]types.Registrant, error)
	GetRegistrant(registrantID int) (types.Registrant, error)
	DeleteRegistrant(tx *sql.Tx, registrantID int) error

	StoreResult(result types.ResultRequest) error
	GetResults(electionID int) ([]types.Result, error)

	Begin() (*sql.Tx, error)
}
