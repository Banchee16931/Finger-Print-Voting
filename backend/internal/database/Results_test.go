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

func TestClient_StoreResult(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name        string
		inputResult types.ResultRequest
		mockExecErr error // Mock the error returned by db.Exec
		expectedErr error
	}{
		{
			name: "happy_path",
			inputResult: types.ResultRequest{
				ElectionID:  1,
				FirstName:   "John",
				LastName:    "Doe",
				Party:       "Independent",
				PartyColour: "#ff000",
				Votes:       100,
			},
			mockExecErr: nil, // No error when executing the query
			expectedErr: nil,
		},
		{
			name: "database_error",
			inputResult: types.ResultRequest{
				ElectionID:  2,
				FirstName:   "Jane",
				LastName:    "Smith",
				Party:       "Democratic",
				PartyColour: "#ff000",
				Votes:       200,
			},
			mockExecErr: errors.New("database error"), // Simulate a database error
			expectedErr: cerr.ErrDB,                   // Expect the cerr.ErrDB error directly
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
			tx, err := db.Begin()
			assert.NoError(t, err, "begin returned an error")

			// Set up expectations for db.Exec
			mock.ExpectExec(`INSERT INTO result (election_id, first_name, last_name, party, party_colour, votes) VALUES ($1, $2, $3, $4, $5, $6);`).
				WithArgs(tc.inputResult.ElectionID, tc.inputResult.FirstName, tc.inputResult.LastName, tc.inputResult.Party, tc.inputResult.PartyColour, tc.inputResult.Votes).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(tc.mockExecErr)

			// Call the function being tested
			err = client.StoreResult(tx, tc.inputResult)

			// Check the returned error
			assert.ErrorIs(t, err, tc.expectedErr, "Incorrect error")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}

func TestClient_GetResults(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name            string
		electionID      int
		mockRows        *sqlmock.Rows
		expectedResults []types.Result
		expectedError   error
	}{
		{
			name:       "happy_path",
			electionID: 1,
			mockRows: sqlmock.NewRows([]string{"result_id", "election_id", "first_name", "last_name", "party", "party_colour", "votes"}).
				AddRow(1, 1, "John", "Doe", "Independent", "#ff000", 100).
				AddRow(2, 1, "Jane", "Smith", "Green", "#ffff0", 200),
			expectedResults: []types.Result{
				{
					ResultID:    1,
					ElectionID:  1,
					FirstName:   "John",
					LastName:    "Doe",
					Party:       "Independent",
					PartyColour: "#ff000",
					Votes:       100,
				},
				{
					ResultID:    2,
					ElectionID:  1,
					FirstName:   "Jane",
					LastName:    "Smith",
					Party:       "Green",
					PartyColour: "#ffff0",
					Votes:       200,
				},
			},
			expectedError: nil,
		},
		{
			name:            "database_error",
			electionID:      3,
			mockRows:        sqlmock.NewRows([]string{"result_id"}).AddRow(1),
			expectedResults: []types.Result{},
			expectedError:   cerr.ErrDB, // Simulate a database error
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
			mock.ExpectQuery(`SELECT result_id, election_id, first_name, last_name, party, party_colour, votes FROM result WHERE election_id=$1;`).
				WithArgs(tc.electionID).
				WillReturnRows(tc.mockRows)

			// Call the function being tested
			results, err := client.GetResults(tc.electionID)

			// Check the returned results and error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")
			assert.Equal(t, results, tc.expectedResults, "Incorrect results")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}
