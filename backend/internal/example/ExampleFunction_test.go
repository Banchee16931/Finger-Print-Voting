package example_test

import (
	"finger-print-voting-backend/internal/example"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowerThanFourty(t *testing.T) {
	// Assign
	input := 39
	expected := true

	// Act
	actual := example.LowerThanFourty(input)

	// Assert
	assert.Equal(t, expected, actual)
}
