package example_test

import (
	"finger-print-voting-backend/internal/example"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowerThanFourty(t *testing.T) {
	// Assign
	expected := 39

	// Act
	actual := example.LowerThanFourty(expected)

	// Assert
	assert.Equal(t, expected, actual)
}

func TestThatFails(t *testing.T) {
	// Assert
	assert.Fail(t, "this one will fail")
}
