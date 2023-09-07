package database_test

import (
	"errors"
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/database"
	"finger-print-voting-backend/internal/types"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestClient_StoreVote(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name        string
		inputVote   types.Vote
		mockExecErr error // Mock the error returned by db.Exec
		expectedErr error
	}{
		{
			name: "happy_path",
			inputVote: types.Vote{
				Username:    "user1",
				ElectionID:  1,
				CandidateID: 100,
			},
			mockExecErr: nil, // No error when executing the query
			expectedErr: nil,
		},
		{
			name: "database_error",
			inputVote: types.Vote{
				Username:    "user2",
				ElectionID:  2,
				CandidateID: 200,
			},
			mockExecErr: errors.New("database error"), // Simulate a database error
			expectedErr: cerr.ErrDB,
		},
	}

	for i := 0; i < len(testCases); i++ {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock DB connection: %v", err)
			}
			defer db.Close()

			// Create a new Client with the mock database connection
			client := database.NewClientFromDatabase(db)

			// creating transaction
			mock.ExpectBegin()
			tx, err := db.Begin()
			assert.NoError(t, err, "failed to begin transaction")

			// Set up expectations for db.Exec
			mock.ExpectExec(`INSERT INTO votes (username, election_id, candidate_id) VALUES ($1, $2, $3);`).
				WithArgs(tc.inputVote.Username, tc.inputVote.ElectionID, tc.inputVote.CandidateID).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tc.mockExecErr)

			// Call the function being tested
			err = client.StoreVote(tx, tc.inputVote)

			// Check the returned errorS
			assert.ErrorIs(t, err, tc.expectedErr, "Incorrect error")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}

func TestClient_GetVotes(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name          string
		electionID    int
		mockRows      *sqlmock.Rows
		expectedVotes []types.Vote
		expectedError error
	}{
		{
			name:       "happy_path",
			electionID: 1,
			mockRows: sqlmock.NewRows([]string{"username", "election_id", "candidate_id"}).
				AddRow("user1", 1, 100).
				AddRow("user2", 1, 200),
			expectedVotes: []types.Vote{
				{
					Username:    "user1",
					ElectionID:  1,
					CandidateID: 100,
				},
				{
					Username:    "user2",
					ElectionID:  1,
					CandidateID: 200,
				},
			},
			expectedError: nil,
		},
		{
			name:          "database_error",
			electionID:    3,
			mockRows:      sqlmock.NewRows([]string{"username"}).AddRow("user3"),
			expectedVotes: []types.Vote{},
			expectedError: cerr.ErrDB, // Simulate a database error
		},
	}

	for i := 0; i < len(testCases); i++ {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock DB connection: %v", err)
			}
			defer db.Close()

			// Create a new Client with the mock database connection
			client := database.NewClientFromDatabase(db)

			// Set up expectations for db.Query
			mock.ExpectQuery(`SELECT username, election_id, candidate_id FROM votes WHERE election_id=$1;`).
				WithArgs(tc.electionID).
				WillReturnRows(tc.mockRows)

			// Call the function being tested
			votes, err := client.GetVotes(tc.electionID)

			// Check the returned votes and error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")
			assert.Equal(t, votes, tc.expectedVotes, "Incorrect votes")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}

func TestClient_DeleteVotes(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name        string
		electionID  int
		mockExecErr error // Mock the error returned by db.Exec
		expectedErr error
	}{
		{
			name:        "happy_path",
			electionID:  1,
			mockExecErr: nil, // No error when executing the query
			expectedErr: nil,
		},
		{
			name:        "database_error",
			electionID:  2,
			mockExecErr: errors.New("database error"), // Simulate a database error
			expectedErr: cerr.ErrDB,
		},
	}

	for i := 0; i < len(testCases); i++ {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock DB connection: %v", err)
			}
			defer db.Close()

			// Create a new Client with the mock database connection
			client := database.NewClientFromDatabase(db)

			// mock transaction begin
			mock.ExpectBegin()
			tx, err := client.Begin()
			assert.NoError(t, err, "begin returned an error")

			// Set up expectations for db.Exec
			mock.ExpectExec(`DELETE FROM votes WHERE election_id=$1;`).
				WithArgs(tc.electionID).
				WillReturnResult(sqlmock.NewResult(0, 1)).
				WillReturnError(tc.mockExecErr)

			// Call the function being tested
			err = client.DeleteVotes(tx, tc.electionID)

			// Check the returned error
			assert.ErrorIs(t, err, tc.expectedErr, "Incorrect error")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}
