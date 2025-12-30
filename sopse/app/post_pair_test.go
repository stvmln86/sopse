package app

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func TestPostPair(t *testing.T) {
	// setup
	app := mockApp(t)
	app.Conf.UserSize = 3

	// success
	app, w := mockRun(t, app, "POST", "/api/mockUser/name", "body")
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "ok", w.Body.String())

	// confirm - database
	body := test.Get(t, app.DB, "select body from Pairs where name='name'")
	assert.Equal(t, "body", body)

	// failure - not found
	_, w = mockRun(t, app, "POST", "/api/nope/name", "body")
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "error 404: user not found", w.Body.String())

	// failure - storage limit reached
	_, w = mockRun(t, app, "POST", "/api/mockUser/name2", "body2")
	assert.Equal(t, http.StatusTooManyRequests, w.Code)
	assert.Equal(t, "error 429: storage limit reached", w.Body.String())
}
