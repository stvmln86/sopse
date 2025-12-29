package app

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIndexOr404(t *testing.T) {
	// success - index page
	_, w := mockRun(t, "GET", "/", "")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())

	// failure - 404 error
	_, w = mockRun(t, "GET", "/nope", "")
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: not found", w.Body.String())
}
