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

	"github.com/stretchr/testify/assert"
)

func TestHandlePostRegistration(t *testing.T) {
	t.Parallel()
	// Assign
	input := types.RegistrationRequest{
		FirstName:       "Test",
		LastName:        "User",
		Email:           "Email",
		PhoneNo:         "420",
		Fingerprint:     "Fingerprint",
		ProofOfIdentity: "Proof",
		Location:        "Location",
	}

	// setting up database expectations
	dbClient := testutils.MockDB{}
	dbClient.On("StoreRegistrant", input).Return(nil)

	srv := api.NewServer()
	srv.WithDBClient(&dbClient)
	srv.WithPasswordSecret("test secret")

	// setting up API inputs
	w := httptest.NewRecorder()
	inputBody, err := json.Marshal(input)
	assert.NoError(t, err, "failed to marshal input")
	r := httptest.NewRequest(http.MethodPost, "/test", bytes.NewReader(inputBody))

	// Act
	srv.HandlePostRegistration(w, r)

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

func TestHandlePostRegistration_InvalidBodyError(t *testing.T) {
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
	srv.HandlePostRegistration(w, r)

	// Assert
	res := w.Result()

	if !assert.Equal(t, http.StatusBadRequest, res.StatusCode, "incorrect status code") {
		return
	}

	dbClient.AssertExpectations(t)

	AssertHTTPErrorResponse(t, res.Body)
}
