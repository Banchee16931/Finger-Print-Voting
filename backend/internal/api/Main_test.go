package api_test

import (
	"encoding/json"
	"finger-print-voting-backend/internal/cerr"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertHTTPErrorResponse(t *testing.T, body io.ReadCloser) bool {
	defer body.Close()
	var commonErr cerr.CommonError
	err := json.NewDecoder(body).Decode(&commonErr)
	if !assert.NoError(t, err, "failed to decode error") {
		outputBody(t, body)
		return false
	}

	if !assert.NotEmpty(t, commonErr.Message, "response doesn't have error message") {
		outputBody(t, body)
		return false
	}

	return true
}

func outputBody(t *testing.T, body io.ReadCloser) {
	readBody, err := io.ReadAll(body)
	assert.NoError(t, err, "failed to read whole body")
	t.Logf("Response Body: %s", string(readBody))
}
