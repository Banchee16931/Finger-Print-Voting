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

func TestClient_StoreRegistrant(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name            string
		registrant      types.RegistrationRequest
		mockExpectation func(mock sqlmock.Sqlmock)
		expectedError   error
	}{
		{
			name: "happy_path",
			registrant: types.RegistrationRequest{
				FirstName:       "John",
				LastName:        "Doe",
				Email:           "john@example.com",
				PhoneNo:         "123-456-7890",
				Fingerprint:     "fingerprint_data",
				ProofOfIdentity: "identity_proof_data",
				Location:        "Hertfordshire",
			},
			mockExpectation: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO registrants (first_name, last_name, email, phone_no, fingerprint, proof, authority_location) VALUES ($1, $2, $3, $4, $5, $6, $7);`).
					WithArgs("John", "Doe", "john@example.com", "123-456-7890", "fingerprint_data", "identity_proof_data", "Hertfordshire").
					WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate a successful insertion
			},
			expectedError: nil,
		},
		{
			name: "database_error",
			registrant: types.RegistrationRequest{
				FirstName:       "Jane",
				LastName:        "Smith",
				Email:           "jane@example.com",
				PhoneNo:         "987-654-3210",
				Fingerprint:     "another_fingerprint",
				ProofOfIdentity: "another_proof_data",
				Location:        "Hertfordshire",
			},
			mockExpectation: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO registrants (first_name, last_name, email, phone_no, fingerprint, proof, authority_location) VALUES ($1, $2, $3, $4, $5, $6, $7);`).
					WithArgs("Jane", "Smith", "jane@example.com", "987-654-3210", "another_fingerprint", "another_proof_data", "Hertfordshire").
					WillReturnError(errors.New("database error")) // Simulate a database error
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

			// Set up expectations for db.Exec
			tc.mockExpectation(mock)

			// Call the function being tested
			err = client.StoreRegistrant(tc.registrant)

			// Check the returned error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}

func TestClient_GetRegistrants(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name            string
		mockRows        *sqlmock.Rows
		expectedResults []types.Registrant
		expectedError   error
	}{
		{
			name: "happy_path",
			mockRows: sqlmock.NewRows([]string{"registrant_id", "first_name", "last_name", "email", "phone_no", "fingerprint", "proof", "authority_location"}).
				AddRow(1, "John", "Doe", "john@example.com", "123-456-7890", "fingerprint_data", "identity_proof_data", "Hertfordshire").
				AddRow(2, "Jane", "Smith", "jane@example.com", "987-654-3210", "another_fingerprint", "another_proof_data", "Hertfordshire"),
			expectedResults: []types.Registrant{
				{
					RegistrantID:    1,
					FirstName:       "John",
					LastName:        "Doe",
					Email:           "john@example.com",
					PhoneNo:         "123-456-7890",
					Fingerprint:     "fingerprint_data",
					ProofOfIdentity: "identity_proof_data",
					Location:        "Hertfordshire",
				},
				{
					RegistrantID:    2,
					FirstName:       "Jane",
					LastName:        "Smith",
					Email:           "jane@example.com",
					PhoneNo:         "987-654-3210",
					Fingerprint:     "another_fingerprint",
					ProofOfIdentity: "another_proof_data",
					Location:        "Hertfordshire",
				},
			},
			expectedError: nil,
		},
		{
			name:            "database_error",
			mockRows:        sqlmock.NewRows([]string{"registrant_id"}).AddRow(1),
			expectedResults: []types.Registrant{},
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
			mock.ExpectQuery(`SELECT registrant_id, first_name, last_name, email, phone_no, fingerprint, proof, authority_location FROM registrants;`).
				WillReturnRows(tc.mockRows)

			// Call the function being tested
			registrants, err := client.GetRegistrants()

			// Check the returned results and error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")
			assert.Equal(t, registrants, tc.expectedResults, "Incorrect results")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}

func TestClient_DeleteRegistrant(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		name            string
		registrantID    int
		mockExpectation func(mock sqlmock.Sqlmock)
		expectedError   error
	}{
		{
			name:         "happy_path",
			registrantID: 1,
			mockExpectation: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM registrants WHERE registrant_id = $1;`).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1)) // Simulate a successful deletion
			},
			expectedError: nil,
		},
		{
			name:         "database_error",
			registrantID: 2,
			mockExpectation: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM registrants WHERE registrant_id = $1;`).
					WithArgs(2).
					WillReturnError(errors.New("database error")) // Simulate a database error
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

			mock.ExpectBegin()
			tx, err := db.Begin()
			assert.NoError(t, err, "failed to begin transaction")

			// Set up expectations for db.Exec
			tc.mockExpectation(mock)

			// Call the function being tested
			err = client.DeleteRegistrant(tx, tc.registrantID)

			// Check the returned error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}
