package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIndexOr404(t *testing.T) {
	// setup
	app := mockApp(t)
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	// success
	app.GetIndexOr404(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, w.Body.String())

	// setup
	r = httptest.NewRequest("GET", "/nope", nil)
	w = httptest.NewRecorder()

	// failure - 404 error
	app.GetIndexOr404(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: not found", w.Body.String())
}
