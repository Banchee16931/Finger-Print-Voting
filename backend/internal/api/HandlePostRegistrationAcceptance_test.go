package api_test

import (
	"bytes"
	"encoding/json"
	"finger-print-voting-backend/internal/api"
	"finger-print-voting-backend/internal/test/testutils"
	"finger-print-voting-backend/internal/types"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandlePostRegistrationAcceptance_Accept(t *testing.T) {
	t.Parallel()
	// Assign
	input := types.UserAcceptanceRequest{
		RegistrantID: 0,
		Accepted:     true,
		Username:     "test_user",
		Password:     "test_password",
	}

	registrationInDB := types.Registrant{
		RegistrantID:    0,
		FirstName:       "Test",
		LastName:        "User",
		Email:           "Email",
		PhoneNo:         "42",
		Fingerprint:     "Fingerprint",
		ProofOfIdentity: "ProofOfIdentity",
		Location:        "Location",
	}

	// setting up database expectations
	dbClient := testutils.MockDB{}
	db, dbMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err, "failed to create mock database")

	dbMock.ExpectBegin()
	dbMock.ExpectCommit()
	tx, err := db.Begin()
	assert.NoError(t, err, "failed to create transaction")

	dbClient.On("GetRegistrant", input.RegistrantID).Return(registrationInDB, nil)
	dbClient.On("Begin").Return(tx, nil)
	dbClient.On("DeleteRegistrant", tx, input.RegistrantID).Return(nil)
	dbClient.On("StoreVoter", tx, mock.Anything).Return(nil)

	srv := api.NewServer()
	srv.WithDBClient(&dbClient)
	srv.WithPasswordSecret("test secret")

	// setting up API inputs
	w := httptest.NewRecorder()
	inputBody, err := json.Marshal(input)
	assert.NoError(t, err, "failed to marshal input")
	r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewReader(inputBody))

	// Act
	srv.HandlePostRegistrationAcceptance(w, r)

	// Assert
	res := w.Result()

	dbClient.AssertExpectations(t)

	// check response status code
	if !assert.Equal(t, http.StatusCreated, res.StatusCode, "incorrect status code") {
		responseBody, err := io.ReadAll(res.Body)
		assert.NoError(t, err, "reading body caused an error")
		t.Logf("Response Body: %s", string(responseBody))
		res.Body.Close()
		return
	}
}

func TestHandlePostRegistrationAcceptance_Decline(t *testing.T) {
	t.Parallel()
	// Assign
	input := types.UserAcceptanceRequest{
		RegistrantID: 0,
		Accepted:     false,
		Username:     "test_user",
		Password:     "test_password",
	}

	registrationInDB := types.Registrant{
		RegistrantID:    0,
		FirstName:       "Test",
		LastName:        "User",
		Email:           "Email",
		PhoneNo:         "42",
		Fingerprint:     "Fingerprint",
		ProofOfIdentity: "ProofOfIdentity",
		Location:        "Location",
	}

	// setting up database expectations
	dbClient := testutils.MockDB{}
	db, dbMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err, "failed to create mock database")

	dbMock.ExpectBegin()
	dbMock.ExpectCommit()
	tx, err := db.Begin()
	assert.NoError(t, err, "failed to create transaction")

	dbClient.On("GetRegistrant", input.RegistrantID).Return(registrationInDB, nil)
	dbClient.On("Begin").Return(tx, nil)
	dbClient.On("DeleteRegistrant", tx, input.RegistrantID).Return(nil)

	srv := api.NewServer()
	srv.WithDBClient(&dbClient)
	srv.WithPasswordSecret("test secret")

	// setting up API inputs
	w := httptest.NewRecorder()
	inputBody, err := json.Marshal(input)
	assert.NoError(t, err, "failed to marshal input")
	r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewReader(inputBody))

	// Act
	srv.HandlePostRegistrationAcceptance(w, r)

	// Assert
	res := w.Result()

	dbClient.AssertExpectations(t)

	// check response status code
	if !assert.Equal(t, http.StatusCreated, res.StatusCode, "incorrect status code") {
		responseBody, err := io.ReadAll(res.Body)
		assert.NoError(t, err, "reading body caused an error")
		t.Logf("Response Body: %s", string(responseBody))
		res.Body.Close()
		return
	}
}

func TestHandlePostRegistrationAcceptance_InvalidBodyError(t *testing.T) {
	t.Parallel()
	// Assign
	invalidJSON := []byte{}

	// setting up database expectations
	dbClient := testutils.MockDB{}

	srv := api.NewServer()
	srv.WithDBClient(&dbClient)
	srv.WithPasswordSecret("test secret")

	// setting up API inputs
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewReader(invalidJSON))

	// Act
	srv.HandlePostRegistrationAcceptance(w, r)

	// Assert
	res := w.Result()

	if !assert.Equal(t, http.StatusBadRequest, res.StatusCode, "incorrect status code") {
		return
	}

	dbClient.AssertExpectations(t)

	AssertHTTPErrorResponse(t, res.Body)
}
