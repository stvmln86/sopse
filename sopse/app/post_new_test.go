package app

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/items/user"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
)

func TestPostNewUser(t *testing.T) {
	// success
	app, w := mockRun(t, "POST", "/api/new", "")
	body := w.Body.String()
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Regexp(t, `[\w-_]{22}`, body)

	// confirm - database
	user, err := user.Get(app.DB, body)
	assert.Equal(t, "192.0.2.1", user.Addr)
	asrt.TimeNow(t, user.Init)
	assert.NoError(t, err)
}
