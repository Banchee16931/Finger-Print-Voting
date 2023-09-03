package api_test

import (
	"encoding/json"
	"finger-print-voting-backend/internal/api"
	"finger-print-voting-backend/internal/api/auth"
	"finger-print-voting-backend/internal/test/testutils"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHandleGetUser(t *testing.T) {
	t.Parallel()
	// Assign
	testUsername := "user"
	testPass := "test_password"
	testSecret := "test secret"
	hashedTestPass, err := bcrypt.GenerateFromPassword([]byte(testPass), bcrypt.MinCost)
	assert.NoError(t, err, "bcrypt returned an error")

	db := testutils.MockDB{}

	expectedUser := types.User{
		Username:  testUsername,
		Password:  string(hashedTestPass),
		Admin:     true,
		FirstName: "Test",
		LastName:  "User",
	}
	db.On("GetUser", testUsername).Return(expectedUser, nil)

	srv := api.NewServer()
	srv.WithDBClient(&db)
	srv.WithPasswordSecret(testSecret)

	jwt, err := auth.GenerateJWT(testSecret, testUsername)
	assert.NoError(t, err, "failed to generate JWT")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Add("Autherization", fmt.Sprintf("Bearer {%s}", jwt))

	srv.MiddlewareAuth(api.AuthLoggedIn)(http.HandlerFunc(srv.HandleGetUser)).ServeHTTP(w, req)

	// Assert
	res := w.Result()
	if !assert.Equal(t, http.StatusOK, res.StatusCode, "incorrect status code") {
		responseBody, err := io.ReadAll(res.Body)
		assert.NoError(t, err, "reading body caused an error")
		t.Logf("Response Body: %s", string(responseBody))
		res.Body.Close()
		return
	}

	body := res.Body
	defer res.Body.Close()

	var userDetails types.UserDetails
	bodyBytes, err := io.ReadAll(body)
	assert.NoError(t, err, "failed to read body")
	assert.NoError(t, json.Unmarshal(bodyBytes, &userDetails), "failed to unmarshal response")
	assert.Equal(t, types.UserDetails{}.FromUser(expectedUser), userDetails, "user did not match expected")
}
