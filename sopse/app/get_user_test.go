package app

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	// success
	_, w := mockRun(t, "GET", "/api/mockUser1", "")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "alpha\nbravo", w.Body.String())

	// failure - not found
	_, w = mockRun(t, "GET", "/api/nope", "body")
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: user not found", w.Body.String())
}
