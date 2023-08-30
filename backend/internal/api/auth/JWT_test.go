package auth_test

import (
	"finger-print-voting-backend/internal/api/auth"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests that a JWT can be geneated and then parsed, maintaining the inputted username
func TestJWT(t *testing.T) {
	testSecret := "my secret"
	testUsername := "test user"
	jwt, err := auth.GenerateJWT(testSecret, testUsername)
	assert.NoError(t, err, "GenerateJWT returned an error")

	claims, err := auth.GetClaims(testSecret, jwt)
	assert.NoError(t, err, "GetClaims returned an error")

	assert.Equal(t, testUsername, claims.Username, "claimed username did not match inputted")
}

// Tests that GetClaims correctly returns an error when the username is blank
func TestJWT_EmptyUsernameError(t *testing.T) {
	testSecret := "my secret"
	testUsername := ""
	jwt, err := auth.GenerateJWT(testSecret, testUsername)
	assert.NoError(t, err, "GenerateJWT returned an error")

	claims, err := auth.GetClaims(testSecret, jwt)
	assert.ErrorIs(t, err, auth.ErrParse, "incorrect error")
	assert.Equal(t, auth.Claims{}, claims, "claims where not blank")
}
