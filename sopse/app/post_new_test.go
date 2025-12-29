package app

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/items/user"
)

func TestPostNewUser(t *testing.T) {
	// success
	app, w := mockRun(t, "POST", "/new", "")
	body := w.Body.String()
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Regexp(t, `[\w-_]{22}`, body)

	// confirm - database
	user, err := user.Get(app.DB, body)
	assert.Equal(t, "192.0.2.1", user.Addr)
	assert.WithinDuration(t, time.Now(), user.Init, 5*time.Second)
	assert.NoError(t, err)
}
