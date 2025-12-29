package app

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/items/pair"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
)

func TestPostPair(t *testing.T) {
	// success
	app, w := mockRun(t, "POST", "/mockUser1/name", "body")
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "ok", w.Body.String())

	// confirm - database
	pair, err := pair.Get(app.DB, "mockUser1", "name")
	assert.Equal(t, "body", pair.Body)
	asrt.TimeNow(t, pair.Init)
	assert.NoError(t, err)

	// failure - not found
	_, w = mockRun(t, "GET", "/nope/name", "body")
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: not found", w.Body.String())
}
