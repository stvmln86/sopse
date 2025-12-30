package app

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/conf"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func mockApp(t *testing.T) *App {
	conf := conf.Parse(nil)
	db := test.DB()
	return New(conf, db)
}

func mockRun(t *testing.T, meth, path, body string) (*App, *httptest.ResponseRecorder) {
	app := mockApp(t)
	b := bytes.NewBufferString(body)
	r := httptest.NewRequest(meth, path, b)
	w := httptest.NewRecorder()
	app.ServeMux().ServeHTTP(w, r)
	return app, w
}

func TestNew(t *testing.T) {
	// setup
	conf := conf.Parse(nil)
	db := test.DB()

	// success
	app := New(conf, db)
	assert.Equal(t, conf, app.Conf)
	assert.Equal(t, db, app.DB)
}

func TestNewParse(t *testing.T) {
	// setup
	elems := []string{"-dbse", ":memory:"}

	// success
	app, err := NewParse(elems)
	assert.NotNil(t, app)
	assert.NoError(t, err)
}

func TestClose(t *testing.T) {
	// setup
	app := mockApp(t)

	// success
	err := app.Close()
	assert.NoError(t, err)

	// confirm - database
	err = app.DB.Ping()
	assert.Error(t, err)
}

func TestServeMux(t *testing.T) {
	// setup
	app := mockApp(t)

	// success
	smux := app.ServeMux()
	assert.NotNil(t, smux)
}

func TestStart(t *testing.T) {
	// cannot test main server
}
