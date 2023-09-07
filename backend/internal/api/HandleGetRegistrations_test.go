package api_test

import (
	"encoding/json"
	"finger-print-voting-backend/internal/api"
	"finger-print-voting-backend/internal/test/testutils"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests that Server.HandleGetRegistrations will correctly return registrations gotten from the database
func TestHandleGetRegistrations(t *testing.T) {
	t.Parallel()
	registrationsInDB := []types.Registrant{{
		RegistrantID:    0,
		FirstName:       "Test",
		LastName:        "User",
		Email:           "Email",
		PhoneNo:         "42",
		Fingerprint:     "Fingerprint",
		ProofOfIdentity: "ProofOfIdentity",
		Location:        "Location",
	}, {
		RegistrantID:    1,
		FirstName:       "Other",
		LastName:        "Guy",
		Email:           "Email2",
		PhoneNo:         "1",
		Fingerprint:     "Ahh",
		ProofOfIdentity: "Test Data",
		Location:        "Arun",
	}}

	// setting up database expectations
	dbClient := testutils.MockDB{}

	dbClient.On("GetRegistrants").Return(registrationsInDB, nil)

	srv := api.NewServer()
	srv.WithDBClient(&dbClient)
	srv.WithPasswordSecret("test secret")

	// setting up API inputs
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", nil)

	// Act
	srv.HandleGetRegistrations(w, r)

	// Assert
	res := w.Result()

	dbClient.AssertExpectations(t)

	// check response status code
	if !assert.Equal(t, http.StatusOK, res.StatusCode, "incorrect status code") {
		responseBody, err := io.ReadAll(res.Body)
		assert.NoError(t, err, "reading body caused an error")
		t.Logf("Response Body: %s", string(responseBody))
		res.Body.Close()
		return
	}

	responseBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	assert.NoError(t, err, "reading body caused an error")
	t.Logf("Response Body: %s", string(responseBody))

	var actualOutputRegistrants []types.Registrant
	assert.NoError(t, json.Unmarshal(responseBody, &actualOutputRegistrants), "failed to read response body")

	assert.ElementsMatch(t, registrationsInDB, actualOutputRegistrants, "response body was not what was expeceted")
}

// Tests that Server.HandleGetRegistrations will correctly return registrations gotten from the database, when there is no registrants in the database
func TestHandleGetRegistrations_NoRegistrations(t *testing.T) {
	t.Parallel()
	registrationsInDB := []types.Registrant{}

	// setting up database expectations
	dbClient := testutils.MockDB{}

	dbClient.On("GetRegistrants").Return(registrationsInDB, nil)

	srv := api.NewServer()
	srv.WithDBClient(&dbClient)
	srv.WithPasswordSecret("test secret")

	// setting up API inputs
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", nil)

	// Act
	srv.HandleGetRegistrations(w, r)

	// Assert
	res := w.Result()

	dbClient.AssertExpectations(t)

	// check response status code
	if !assert.Equal(t, http.StatusOK, res.StatusCode, "incorrect status code") {
		responseBody, err := io.ReadAll(res.Body)
		assert.NoError(t, err, "reading body caused an error")
		t.Logf("Response Body: %s", string(responseBody))
		res.Body.Close()
		return
	}

	var body []byte
	_, err := res.Body.Read(body)
	assert.NoError(t, err, "failed to read response body")

	defer res.Body.Close()

	assert.Equal(t, "", string(body), "response body was not what was expeceted")
}

// Tests that if the database runs into an error it is correctly reported by HandleGetRegistrations
func TestHandleGetRegistrations_DBError(t *testing.T) {
	t.Parallel()

	// setting up database expectations
	dbClient := testutils.MockDB{}

	dbClient.On("GetRegistrants").Return([]types.Registrant{}, fmt.Errorf("I am an error"))

	srv := api.NewServer()
	srv.WithDBClient(&dbClient)
	srv.WithPasswordSecret("test secret")

	// setting up API inputs
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/test", nil)

	// Act
	srv.HandleGetRegistrations(w, r)

	// Assert
	dbClient.AssertExpectations(t)

	res := w.Result()
	if !assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "incorrect status code") {
		return
	}

	AssertHTTPErrorResponse(t, res.Body)
}
