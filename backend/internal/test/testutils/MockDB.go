package testutils

import (
	"database/sql"
	"finger-print-voting-backend/internal/types"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (client *MockDB) StoreElection(election types.ElectionRequest) error {
	call := client.Called(election)
	return call.Error(0)
}

func (client *MockDB) GetElections() ([]types.Election, error) {
	call := client.Called()
	return call.Get(0).([]types.Election), call.Error(1)
}

func (client *MockDB) GetElectionByLocation(location string) ([]types.Election, error) {
	call := client.Called(location)
	return call.Get(0).([]types.Election), call.Error(1)
}

func (client *MockDB) GetCandidates(electionID int) ([]types.Candidate, error) {
	call := client.Called(electionID)
	return call.Get(0).([]types.Candidate), call.Error(1)
}

func (client *MockDB) DeleteCandidates(electionID int) error {
	call := client.Called(electionID)
	return call.Error(0)
}

func (client *MockDB) StoreRegistrant(registrant types.RegistrationRequest) error {
	call := client.Called(registrant)
	return call.Error(0)
}

func (client *MockDB) GetRegistrants() ([]types.Registrant, error) {
	call := client.Called()
	return call.Get(0).([]types.Registrant), call.Error(1)
}

func (client *MockDB) GetRegistrant(registrantID int) (types.Registrant, error) {
	call := client.Called(registrantID)
	return call.Get(0).(types.Registrant), call.Error(1)
}

func (client *MockDB) DeleteRegistrant(tx *sql.Tx, registrantID int) error {
	call := client.Called(tx, registrantID)
	return call.Error(0)
}

func (client *MockDB) StoreResult(result types.ResultRequest) error {
	call := client.Called(result)
	return call.Error(0)
}

func (client *MockDB) GetResults(electionID int) ([]types.Result, error) {
	call := client.Called(electionID)
	return call.Get(0).([]types.Result), call.Error(1)
}

func (client *MockDB) IsSchemaSetup() (bool, error) {
	call := client.Called()
	return call.Bool(0), call.Error(1)
}

func (client *MockDB) SetupSchema() error {
	call := client.Called()
	return call.Error(0)
}

func (client *MockDB) StoreUser(tx *sql.Tx, user types.User) error {
	call := client.Called(tx, user)
	return call.Error(0)
}

func (client *MockDB) StoreVoter(tx *sql.Tx, voter types.Voter) error {
	call := client.Called(tx, voter)
	return call.Error(0)
}

func (client *MockDB) GetVoter(username string) (types.Voter, error) {
	call := client.Called(username)
	return call.Get(0).(types.Voter), call.Error(1)
}

func (client *MockDB) GetUser(username string) (types.User, error) {
	call := client.Called(username)
	return call.Get(0).(types.User), call.Error(1)
}

func (client *MockDB) DeleteVoter(username string) error {
	call := client.Called(username)
	return call.Error(0)
}

func (client *MockDB) StoreVote(vote types.Vote) error {
	call := client.Called(vote)
	return call.Error(0)
}

func (client *MockDB) GetVotes(electionID int) ([]types.Vote, error) {
	call := client.Called(electionID)
	return call.Get(0).([]types.Vote), call.Error(1)
}

func (client *MockDB) DeleteVotes(electionID int) error {
	call := client.Called(electionID)
	return call.Error(0)
}

func (client *MockDB) Begin() (*sql.Tx, error) {
	call := client.Called()
	return call.Get(0).(*sql.Tx), call.Error(1)
}
