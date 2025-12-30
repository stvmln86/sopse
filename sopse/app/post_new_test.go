package app

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func TestPostNewUser(t *testing.T) {
	// success
	app, w := mockRun(t, "POST", "/api/new", "")
	body := w.Body.String()
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Regexp(t, `\w{32}`, body)

	// confirm - database
	addr := test.Get(t, app.DB, "select addr from Users where uuid=?", body)
	assert.Equal(t, "192.0.2.1", addr)
}
