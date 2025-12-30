package app

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIndexOr404(t *testing.T) {
	// setup
	app := mockApp(t)

	// success - index page
	_, w := mockRun(t, app, "GET", "/", "")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())

	// failure - not found
	_, w = mockRun(t, app, "GET", "/nope", "")
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: url not found", w.Body.String())
}
