package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIndex(t *testing.T) {
	// setup
	app := mockApp(t)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	// success
	app.GetIndex(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, Index, w.Body.String())

	// setup
	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/nope", nil)

	// failure - 404 error
	app.GetIndex(w, r)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: path /nope not found", w.Body.String())
}
