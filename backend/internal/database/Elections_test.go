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

func TestClient_StoreElection(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name             string
		electionRequest  types.ElectionRequest
		mockExpectations func(mock sqlmock.Sqlmock)
		expectedError    error
		expectedRollback bool
	}{
		{
			name: "happy_path",
			electionRequest: types.ElectionRequest{
				Start:    "2023-09-01",
				End:      "2023-09-02",
				Location: "Hertfordshire",
				Candidates: []types.CandidateRequest{
					{FirstName: "John", LastName: "Doe", Party: "Independent", PartyColour: "Grey", Photo: "john.jpg"},
					{FirstName: "Jane", LastName: "Smith", Party: "Green", PartyColour: "Green", Photo: "jane.png"},
				},
			},
			mockExpectations: func(mock sqlmock.Sqlmock) {
				// Expect a Begin transaction call
				mock.ExpectBegin()

				// Expect an INSERT INTO elections query and returning the election_id
				mock.ExpectQuery(`INSERT INTO elections (election_start, election_end, authority_location) VALUES ($1, $2, $3) RETURNING election_id;`).
					WithArgs("2023-09-01", "2023-09-02", "Hertfordshire").
					WillReturnRows(sqlmock.NewRows([]string{"election_id"}).AddRow(1))

				// Expect StoreCandidates to be called
				mock.ExpectExec(`INSERT INTO candidates (election_id, first_name, last_name, party, party_colour, photo) VALUES ($1, $2, $3, $4, $5, $6);`).
					WithArgs(1, "John", "Doe", "Independent", "Grey", "john.jpg").
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectExec(`INSERT INTO candidates (election_id, first_name, last_name, party, party_colour, photo) VALUES ($1, $2, $3, $4, $5, $6);`).
					WithArgs(1, "Jane", "Smith", "Green", "Green", "jane.png").
					WillReturnResult(sqlmock.NewResult(1, 1))

				// Expect a Commit transaction call
				mock.ExpectCommit()
			},
			expectedError:    nil,
			expectedRollback: false,
		},
		{
			name: "database_error",
			electionRequest: types.ElectionRequest{
				Start:    "2023-09-01",
				End:      "2023-09-02",
				Location: "Hertfordshire",
				Candidates: []types.CandidateRequest{
					{FirstName: "Bob", LastName: "Johnson", Party: "Labour", PartyColour: "Red", Photo: "bob.jpg"},
					{FirstName: "Alice", LastName: "Brown", Party: "Conservative", PartyColour: "Blue", Photo: "jane.png"},
				},
			},
			mockExpectations: func(mock sqlmock.Sqlmock) {
				// Expect a Begin transaction call
				mock.ExpectBegin()

				// Expect an INSERT INTO elections query to return an error
				mock.ExpectQuery(`INSERT INTO elections (election_start, election_end, authority_location) VALUES ($1, $2, $3) RETURNING election_id;`).
					WithArgs("2023-09-01", "2023-09-02", "Hertfordshire").
					WillReturnError(errors.New("database error"))

				// Expect a Rollback transaction call
				mock.ExpectRollback()
			},
			expectedError:    cerr.ErrDB,
			expectedRollback: true,
		},
		// Add more test cases for other scenarios as needed
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

			tc.mockExpectations(mock)

			// Call the function being tested
			err = client.StoreElection(tc.electionRequest)

			// Check the returned error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")

			if tc.expectedRollback {
				// If a rollback is expected, ensure it occurred
				assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
			} else {
				// If no rollback is expected, ensure Commit was called and all expectations were met
				assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
			}
		})
	}
}

func TestClient_GetElections(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name              string
		mockRows          *sqlmock.Rows
		expectedElections []types.Election
		expectedError     error
	}{
		{
			name: "happy_path",
			mockRows: sqlmock.NewRows([]string{"election_id", "election_start", "election_end", "authority_location"}).
				AddRow(1, "2023-09-01", "2023-09-02", "Hertfordshire").
				AddRow(2, "2023-09-03", "2023-09-04", "Hertfordshire"),
			expectedElections: []types.Election{
				{
					ElectionID: 1,
					Start:      "2023-09-01",
					End:        "2023-09-02",
					Location:   "Hertfordshire",
				},
				{
					ElectionID: 2,
					Start:      "2023-09-03",
					End:        "2023-09-04",
					Location:   "Hertfordshire",
				},
			},
			expectedError: nil,
		},
		{
			name:              "database_error",
			mockRows:          sqlmock.NewRows([]string{"election_id"}).AddRow("1"),
			expectedElections: []types.Election{},
			expectedError:     cerr.ErrDB,
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
			mock.ExpectQuery(`SELECT election_id, election_start, election_end, authority_location FROM elections;`).
				WillReturnRows(tc.mockRows)

			// Call the function being tested
			elections, err := client.GetElections()

			// Check the returned results and error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")
			assert.Equal(t, elections, tc.expectedElections, "Incorrect elections")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}

func TestClient_GetCandidates(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name               string
		mockRows           *sqlmock.Rows
		expectedCandidates []types.Candidate
		expectedError      error
	}{
		{
			name: "happy_path",
			mockRows: sqlmock.NewRows([]string{"candidate_id", "election_id", "first_name", "last_name", "party", "party_colour", "photo"}).
				AddRow(1, 1, "John", "Doe", "Independent", "Green", "photo1").
				AddRow(2, 1, "Jane", "Smith", "Green", "Blue", "photo2"),
			expectedCandidates: []types.Candidate{
				{
					CandidateID: 1,
					ElectionID:  1,
					FirstName:   "John",
					LastName:    "Doe",
					Party:       "Independent",
					PartyColour: "Green",
					Photo:       "photo1",
				},
				{
					CandidateID: 2,
					ElectionID:  1,
					FirstName:   "Jane",
					LastName:    "Smith",
					Party:       "Green",
					PartyColour: "Blue",
					Photo:       "photo2",
				},
			},
			expectedError: nil,
		},
		{
			name:               "database_error",
			mockRows:           sqlmock.NewRows([]string{"candidate_id"}).AddRow("1"),
			expectedCandidates: []types.Candidate{},
			expectedError:      cerr.ErrDB,
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
			mock.ExpectQuery(`SELECT candidate_id, election_id, first_name, last_name, party, party_colour, photo FROM candidates WHERE election_id=$1;`).
				WithArgs(1).
				WillReturnRows(tc.mockRows)

			// Call the function being tested
			candidates, err := client.GetCandidates(1) // Pass a dummy election ID

			// Check the returned results and error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")
			assert.Equal(t, candidates, tc.expectedCandidates, "Incorrect candidates")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}

func TestClient_DeleteCandidates(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name            string
		electionID      int
		mockExpectation func(mock sqlmock.Sqlmock)
		expectedError   error
	}{
		{
			name:       "happy_path",
			electionID: 1,
			mockExpectation: func(mock sqlmock.Sqlmock) {
				// Expect a DELETE FROM candidates query
				mock.ExpectExec(`DELETE FROM candidates WHERE election_id = $1;`).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1)) // Simulate a successful deletion
			},
			expectedError: nil,
		},
		{
			name:       "database_error",
			electionID: 2,
			mockExpectation: func(mock sqlmock.Sqlmock) {
				// Expect a DELETE FROM candidates query to return an error
				mock.ExpectExec(`DELETE FROM candidates WHERE election_id = $1;`).
					WithArgs(2).
					WillReturnError(errors.New("database error"))
			},
			expectedError: cerr.ErrDB,
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
			tc.mockExpectation(mock)

			// Call the function being tested
			err = client.DeleteCandidates(tx, tc.electionID)

			// Check the returned error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}
