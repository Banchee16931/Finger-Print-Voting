package database_test

import (
	"database/sql"
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/database"
	"finger-print-voting-backend/internal/types"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetUser(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		username      string
		mockRows      *sqlmock.Rows
		expectedUser  types.User
		expectedError error
	}{
		{
			username: "existing_user",
			mockRows: sqlmock.NewRows([]string{"username", "encrypted_password", "is_admin", "first_name", "last_name"}).
				AddRow("existing_user", "hashed_password", false, "John", "Doe"),
			expectedUser: types.User{
				Username:  "existing_user",
				Password:  "hashed_password",
				Admin:     false,
				FirstName: "John",
				LastName:  "Doe",
			},
			expectedError: nil,
		},
		{
			username:      "non_existing_user",
			mockRows:      sqlmock.NewRows([]string{}), // No rows returned, simulating sql.ErrNoRows
			expectedUser:  types.User{},
			expectedError: cerr.ErrNotFound,
		},
		{
			username:      "error_user",
			mockRows:      sqlmock.NewRows([]string{"username"}).AddRow("error_user"),
			expectedUser:  types.User{},
			expectedError: cerr.ErrDB, // Simulate a database error
		},
	}

	for i := 0; i < len(testCases); i++ {
		tc := testCases[i]

		t.Run(tc.username, func(t *testing.T) {
			t.Parallel()

			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock DB connection: %v", err)
			}
			defer db.Close()

			// Create a new Client with the mock database connection
			client := database.NewClientFromDatabase(db)

			// Set up expectations
			mock.ExpectQuery(`SELECT username, encrypted_password, is_admin, first_name, last_name FROM users WHERE username=$1;`).
				WithArgs(tc.username).
				WillReturnRows(tc.mockRows)

			// Call the function being tested
			user, err := client.GetUser(tc.username)

			// Check the returned user and error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")
			assert.Equal(t, user, tc.expectedUser, "Incorrect user")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}

func TestClient_GetVoter(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		username        string
		mockRows        *sqlmock.Rows
		expectedVoter   types.Voter
		expectedError   error
		expectedUserErr error
	}{
		{
			username: "existing_user",
			mockRows: sqlmock.NewRows([]string{"phone_no", "email", "fingerprint", "authority_location"}).
				AddRow("1234567890", "user@example.com", "fingerprint123", "Location A"),
			expectedVoter: types.Voter{
				PhoneNo:     "1234567890",
				Email:       "user@example.com",
				Fingerprint: "fingerprint123",
				Location:    "Location A",
				User: types.User{
					Username:  "existing_user",
					Password:  "hashed_password", // Assuming you have the same password for this user as in the previous test
					Admin:     false,
					FirstName: "John",
					LastName:  "Doe",
				},
			},
			expectedError:   nil,
			expectedUserErr: nil,
		},
		{
			username: "error_user",
			mockRows: sqlmock.NewRows([]string{"phone_no", "email", "fingerprint", "authority_location"}).
				AddRow("1234567890", "user@example.com", "fingerprint123", "Location A"),
			expectedVoter:   types.Voter{},
			expectedError:   cerr.ErrDB, // Simulate a database error for GetVoter
			expectedUserErr: cerr.ErrDB, // Simulate a database error for GetUser
		},
	}

	for i := 0; i < len(testCases); i++ {
		tc := testCases[i]

		t.Run(tc.username, func(t *testing.T) {
			t.Parallel()

			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock DB connection: %v", err)
			}
			defer db.Close()

			// Create a new Client with the mock database connection
			client := database.NewClientFromDatabase(db)

			// Set up expectations for GetVoter
			mock.ExpectQuery(`SELECT phone_no, email, fingerprint, authority_location FROM voter_details WHERE username=$1;`).
				WithArgs(tc.username).
				WillReturnRows(tc.mockRows)

			// Set up expectations for GetUser
			if tc.expectedUserErr != nil {
				mock.ExpectQuery(`SELECT username, encrypted_password, is_admin, first_name, last_name FROM users WHERE username=$1;`).
					WithArgs(tc.username).
					WillReturnError(tc.expectedUserErr)
			} else {
				mock.ExpectQuery(`SELECT username, encrypted_password, is_admin, first_name, last_name FROM users WHERE username=$1;`).
					WithArgs(tc.username).
					WillReturnRows(sqlmock.NewRows([]string{"username", "encrypted_password", "is_admin", "first_name", "last_name"}).
						AddRow(tc.username, "hashed_password", false, "John", "Doe"))
			}

			// Call the function being tested
			voter, err := client.GetVoter(tc.username)

			// Check the returned voter and error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")
			assert.Equal(t, voter, tc.expectedVoter, "Incorrect voter")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}

func TestClient_GetVoter_NonExistingUserError(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		username        string
		mockRows        *sqlmock.Rows
		expectedVoter   types.Voter
		expectedError   error
		expectedUserErr error
	}{
		{
			username:        "non_existing_user",
			mockRows:        sqlmock.NewRows([]string{}), // No rows returned, simulating sql.ErrNoRows
			expectedVoter:   types.Voter{},
			expectedError:   cerr.ErrNotFound,
			expectedUserErr: cerr.ErrNotFound,
		},
	}

	for i := 0; i < len(testCases); i++ {
		tc := testCases[i]

		t.Run(tc.username, func(t *testing.T) {
			t.Parallel()

			// Create a new mock database connection
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("Failed to create mock DB connection: %v", err)
			}
			defer db.Close()

			// Create a new Client with the mock database connection
			client := database.NewClientFromDatabase(db)

			// Set up expectations for GetVoter
			mock.ExpectQuery(`SELECT phone_no, email, fingerprint, authority_location FROM voter_details WHERE username=$1;`).
				WithArgs(tc.username).
				WillReturnRows(tc.mockRows)

			// Call the function being tested
			voter, err := client.GetVoter(tc.username)

			// Check the returned voter and error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")
			assert.Equal(t, voter, tc.expectedVoter, "Incorrect voter")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}

func TestClient_StoreUser(t *testing.T) {
	t.Parallel()

	// Define test cases
	testCases := []struct {
		user           types.User
		mockExecResult sql.Result
		expectedError  error
	}{
		{
			user: types.User{
				Username:  "new_user",
				Password:  "hashed_password",
				Admin:     false,
				FirstName: "Alice",
				LastName:  "Johnson",
			},
			mockExecResult: sqlmock.NewResult(1, 1), // Simulate a successful INSERT
			expectedError:  nil,
		},
		{
			user: types.User{
				Username:  "error_user",
				Password:  "hashed_password",
				Admin:     false,
				FirstName: "Error",
				LastName:  "User",
			},
			mockExecResult: nil, // Simulate a database error
			expectedError:  cerr.ErrDB,
		},
	}

	for i := 0; i < len(testCases); i++ {
		tc := testCases[i]

		t.Run(tc.user.Username, func(t *testing.T) {
			t.Parallel()

			// Create a new mock database connection
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock DB connection: %v", err)
			}
			defer db.Close()

			// Create a new Client with the mock database connection
			client := database.NewClientFromDatabase(db)
			// Set up expectations
			mock.ExpectExec(`INSERT INTO users \(username, encrypted_password, is_admin, first_name, last_name\) VALUES \(\$1, \$2, \$3, \$4, \$5\);`).
				WithArgs(tc.user.Username, tc.user.Password, tc.user.Admin, tc.user.FirstName, tc.user.LastName).
				WillReturnResult(tc.mockExecResult)

			mock.ExpectBegin()
			tx, err := db.Begin()

			// Call the function being tested
			err = client.StoreUser(tx, tc.user)

			// Check the returned error
			assert.ErrorIs(t, err, tc.expectedError, "Incorrect error type")

			// Ensure all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet(), "There were unfulfilled expectations")
		})
	}
}
