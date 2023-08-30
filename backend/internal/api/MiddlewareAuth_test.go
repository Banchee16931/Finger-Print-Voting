package api_test

import (
	"finger-print-voting-backend/internal/api"
	"finger-print-voting-backend/internal/api/auth"
	"finger-print-voting-backend/internal/test/testutils"
	"finger-print-voting-backend/internal/types"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestMiddlewareAuth(t *testing.T) {
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

	nextHandler := &testutils.MockNextHandler{}

	// Act
	srv.MiddlewareAuth(nextHandler).ServeHTTP(w, req)

	// Assert
	userCtx := nextHandler.Request(t).Context().Value(types.UserContext)

	userDetails, ok := userCtx.(types.UserDetails)
	assert.True(t, ok, "failed to converted user context")
	assert.Equal(t, types.UserDetails{}.FromUser(expectedUser), userDetails, "user details do not match expected")
}

func TestMiddlewareAuth_MissingAuthError(t *testing.T) {
	t.Parallel()
	// Assign
	testSecret := "test secret"

	db := testutils.MockDB{}

	srv := api.NewServer()
	srv.WithDBClient(&db)
	srv.WithPasswordSecret(testSecret)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)

	nextHandler := &testutils.MockNextHandler{}

	// Act
	srv.MiddlewareAuth(nextHandler).ServeHTTP(w, req)

	// Assert
	nextHandler.NotCalled(t)

	result := w.Result()
	assert.Equal(t, http.StatusUnauthorized, result.StatusCode, "incorrect status code")
}

func TestMiddlewareAuth_IncorrectAuthFormatError(t *testing.T) {
	t.Parallel()
	// Assign
	testUsername := "user"
	testSecret := "test secret"

	cases := []struct {
		key   string
		value string
	}{
		{key: "", value: "Bearer {%s}"},
		{key: "Autherisation", value: "Bearer {%s}"},
		{key: "Autherisation", value: "Bearer {%s}"},
		{key: "{Autherisation", value: "Bearer {%s}"},
		{key: "AutherisationD", value: "Bearer {%s}"},

		{key: "Autherization", value: ""},
		{key: "Autherization", value: "Bearer{%s}"},
		{key: "Autherization", value: "Beaer {%s}"},
		{key: "Autherization", value: "Bearer %s}"},
		{key: "Autherization", value: "Bearer {%s"},
		{key: "Autherization", value: "%s}"},
		{key: "Autherization", value: "Bearer {%s}}"},
		{key: "Autherization", value: "{Bearer {%s}"},
	}

	for i := 0; i < len(cases); i++ {
		c := cases[i]

		t.Run(c.key+":"+c.value, func(t *testing.T) {
			t.Parallel()
			db := testutils.MockDB{}

			srv := api.NewServer()
			srv.WithDBClient(&db)
			srv.WithPasswordSecret(testSecret)

			jwt, err := auth.GenerateJWT(testSecret, testUsername)
			assert.NoError(t, err, "failed to generate JWT")

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			req.Header.Add(c.key, fmt.Sprintf(c.value, jwt))

			nextHandler := &testutils.MockNextHandler{}

			// Act
			srv.MiddlewareAuth(nextHandler).ServeHTTP(w, req)

			// Assert
			nextHandler.NotCalled(t)

			result := w.Result()
			assert.Equal(t, http.StatusUnauthorized, result.StatusCode, "incorrect status code")
		})
	}
}

func TestMiddlewareAuth_NoKeyError(t *testing.T) {
	t.Parallel()
	// Assign
	testSecret := "test secret"

	db := testutils.MockDB{}

	srv := api.NewServer()
	srv.WithDBClient(&db)
	srv.WithPasswordSecret(testSecret)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Add("Autherization", "Bearer {}")

	nextHandler := &testutils.MockNextHandler{}

	// Act
	srv.MiddlewareAuth(nextHandler).ServeHTTP(w, req)

	// Assert
	nextHandler.NotCalled(t)

	result := w.Result()
	assert.Equal(t, http.StatusUnauthorized, result.StatusCode, "incorrect status code")
}

func TestMiddlewareAuth_NotInDBError(t *testing.T) {
	t.Parallel()
	// Assign
	testUsername := "user"
	testSecret := "test secret"

	db := testutils.MockDB{}

	db.On("GetUser", testUsername).Return(types.User{}, fmt.Errorf("test error"))

	srv := api.NewServer()
	srv.WithDBClient(&db)
	srv.WithPasswordSecret(testSecret)

	jwt, err := auth.GenerateJWT(testSecret, testUsername)
	assert.NoError(t, err, "failed to generate JWT")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Add("Autherization", fmt.Sprintf("Bearer {%s}", jwt))

	nextHandler := &testutils.MockNextHandler{}

	// Act
	srv.MiddlewareAuth(nextHandler).ServeHTTP(w, req)

	// Assert
	nextHandler.NotCalled(t)

	result := w.Result()
	assert.Equal(t, http.StatusUnauthorized, result.StatusCode, "incorrect status code")
}
