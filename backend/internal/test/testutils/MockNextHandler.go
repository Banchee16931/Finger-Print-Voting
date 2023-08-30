package testutils

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockNextHandler struct {
	request *http.Request
	called  bool
}

func (h *MockNextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write(nextHandlerResponseBody)
	if err != nil {
		panic("failed to write mock next handler body")
	}

	h.request = r
	h.called = true

	w.WriteHeader(http.StatusOK)
}

func (h *MockNextHandler) Request(t *testing.T) *http.Request {
	t.Helper()

	if !h.called {
		assert.FailNow(t, "returned request is not valid as the mocked handler was never called")
	}

	if h.request == nil {
		assert.FailNow(t, "request was never passed to MockNextHandler")
	}

	return h.request
}

func (h *MockNextHandler) Called(t *testing.T) bool {
	t.Helper()

	called := h.called

	assert.True(t, called, "MockNextHandler wasn't called")

	return called
}

func (h *MockNextHandler) NotCalled(t *testing.T) bool {
	t.Helper()

	called := h.called

	assert.False(t, called, "MockNextHandler was called")

	return !called
}

var nextHandlerResponseBody = []byte("A MockNextHandler has served this")
