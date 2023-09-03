package api_test

import (
	"finger-print-voting-backend/internal/api"
	"finger-print-voting-backend/internal/test/testutils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	t.Parallel()

	// Assign
	routes := []struct {
		Method string
		Path   string
	}{
		{
			Method: http.MethodPost,
			Path:   "/login",
		},
		{
			Method: http.MethodGet,
			Path:   "/users",
		},
	}

	db := testutils.MockDB{}

	srv := api.NewServer().WithDBClient(&db).WithPasswordSecret("test secret")
	srv.Start(httptest.DefaultRemoteAddr)

	for i := 0; i < len(routes); i++ {
		route := routes[i]

		t.Run(route.Method+">"+strings.ReplaceAll(strings.Trim(route.Path, "/"), "/", "\\"), func(t *testing.T) {
			t.Parallel()
			// Assign
			req := httptest.NewRequest(route.Method, route.Path, nil)
			w := httptest.NewRecorder()

			// Act
			srv.ServeHTTP(w, req)

			// Assert
			res := w.Result()
			defer res.Body.Close()

			assert.NotEqual(t, http.StatusNotFound, res.StatusCode, "returned not found status code")
		})
	}

	assert.NoError(t, srv.Close(), "failed to close server")
}
