package api_test

import (
	"bytes"
	"encoding/json"
	"finger-print-voting-backend/internal/api"
	"finger-print-voting-backend/internal/cerr"
	"finger-print-voting-backend/internal/test/testutils"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHandlePostLogin(t *testing.T) {
	t.Parallel()
	// Assign
	testUsername := "user"
	testPass := "test_password"
	testSecret := "test secret"
	hashedTestPass, err := bcrypt.GenerateFromPassword([]byte(testPass), bcrypt.MinCost)
	assert.NoError(t, err, "bcrypt returned an error")

	db := testutils.MockDB{}

	db.On("GetUser", testUsername).Return(types.User{
		Username:  testUsername,
		Password:  string(hashedTestPass),
		Admin:     true,
		FirstName: "Test",
		LastName:  "User",
	}, nil)

	srv := api.NewServer()
	srv.WithDBClient(&db)
	srv.WithPasswordSecret(testSecret)

	input := types.LoginRequest{
		Username: testUsername,
		Password: testPass,
	}

	inputBody, err := json.Marshal(input)
	assert.NoError(t, err, "failed to marshal input")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader(inputBody))
	srv.HandlePostLogin(w, req)

	// Assert
	res := w.Result()

	if !assert.Equal(t, http.StatusOK, res.StatusCode, "incorrect status code") {
		responseBody, err := io.ReadAll(res.Body)
		assert.NoError(t, err, "reading body caused an error")
		t.Logf("Response Body: %s", string(responseBody))
		res.Body.Close()
		return
	}

	sessionSet := w.Header().Get("Set-Cookie")
	assert.NotEmpty(t, sessionSet, "set cookie header was empty")
	t.Log(sessionSet)

	split := strings.Split(sessionSet, ":")
	assert.Len(t, split, 2, "split session is not correct length")
	assert.Equal(t, "Autherization", strings.TrimSpace(split[0]))
	assert.Regexp(t, regexp.MustCompile("^(?i)Bearer {(.+)(?-i)}"), strings.TrimSpace(split[1]), "bearer was not in correct format")
}

func TestHandlePostLogin_InvalidRequestBody(t *testing.T) {
	t.Parallel()
	// Assign
	testSecret := "test secret"

	db := testutils.MockDB{}

	srv := api.NewServer()
	srv.WithDBClient(&db)
	srv.WithPasswordSecret(testSecret)

	input := struct {
		Thisisinvalid string `json:"thisisinvalid"`
	}{
		Thisisinvalid: "invalid",
	}

	inputBody, err := json.Marshal(input)
	assert.NoError(t, err, "failed to marshal input")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader(inputBody))
	srv.HandlePostLogin(w, req)

	// Assert
	res := w.Result()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode, "incorrect status code")

	sessionSet := w.Header().Get("Set-Cookie")
	assert.Empty(t, sessionSet, "set cookie header was not empty")
}

func TestHandlePostLogin_UserNotInDB(t *testing.T) {
	t.Parallel()
	// Assign
	testUsername := "user"
	testPass := "test_password"
	testSecret := "test secret"

	db := testutils.MockDB{}

	db.On("GetUser", testUsername).Return(types.User{}, cerr.ErrNotFound)

	srv := api.NewServer()
	srv.WithDBClient(&db)
	srv.WithPasswordSecret(testSecret)

	input := types.LoginRequest{
		Username: testUsername,
		Password: testPass,
	}

	inputBody, err := json.Marshal(input)
	assert.NoError(t, err, "failed to marshal input")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader(inputBody))
	srv.HandlePostLogin(w, req)

	// Assert
	res := w.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "incorrect status code")

	sessionSet := w.Header().Get("Set-Cookie")
	assert.Empty(t, sessionSet, "set cookie header was not empty")
}

func TestHandlePostLogin_IncorrectPassword(t *testing.T) {
	t.Parallel()
	// Assign
	testUsername := "user"
	testPass := "test_password"
	testSecret := "test secret"

	db := testutils.MockDB{}

	db.On("GetUser", testUsername).Return(types.User{
		Username:  testUsername,
		Password:  "bad pass",
		Admin:     true,
		FirstName: "Test",
		LastName:  "User",
	}, nil)

	srv := api.NewServer()
	srv.WithDBClient(&db)
	srv.WithPasswordSecret(testSecret)

	input := types.LoginRequest{
		Username: testUsername,
		Password: testPass,
	}

	inputBody, err := json.Marshal(input)
	assert.NoError(t, err, "failed to marshal input")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader(inputBody))
	srv.HandlePostLogin(w, req)

	// Assert
	res := w.Result()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode, "incorrect status code")

	sessionSet := w.Header().Get("Set-Cookie")
	assert.Empty(t, sessionSet, "set cookie header was not empty")
}

func TestHandlePostLogin_DBError(t *testing.T) {
	t.Parallel()
	// Assign
	testUsername := "user"
	testPass := "test_password"
	testSecret := "test secret"

	db := testutils.MockDB{}

	db.On("GetUser", testUsername).Return(types.User{}, fmt.Errorf("DB Error"))

	srv := api.NewServer()
	srv.WithDBClient(&db)
	srv.WithPasswordSecret(testSecret)

	input := types.LoginRequest{
		Username: testUsername,
		Password: testPass,
	}

	inputBody, err := json.Marshal(input)
	assert.NoError(t, err, "failed to marshal input")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", bytes.NewReader(inputBody))
	srv.HandlePostLogin(w, req)

	// Assert
	res := w.Result()

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode, "incorrect status code")

	sessionSet := w.Header().Get("Set-Cookie")
	assert.Empty(t, sessionSet, "set cookie header was not empty")
}
