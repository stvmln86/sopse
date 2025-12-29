package app

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/mock"
	"github.com/stvmln86/sopse/sopse/tools/conf"
)

func mockApp(t *testing.T) *App {
	conf := conf.Parse(nil)
	db := mock.DB(t)
	return New(conf, db)
}

func TestNew(t *testing.T) {
	// setup
	conf := conf.Parse(nil)
	db := mock.DB(t)

	// success
	app := New(conf, db)
	assert.Equal(t, conf, app.Conf)
	assert.Equal(t, db, app.DB)
}

func TestNewParse(t *testing.T) {
	// setup
	dire := t.TempDir()
	path := filepath.Join(dire, t.Name()+".db")
	elems := []string{"-dbse", path}

	// success
	app, err := NewParse(elems)
	assert.NoError(t, err)

	// confirm - database
	assert.Contains(t, app.DB.Path(), "TestNewParse.db")
}

func TestClose(t *testing.T) {
	// setup
	app := mockApp(t)

	// success
	err := app.Close()
	assert.NoError(t, err)

	// confirm - database
	_, err = app.DB.Begin(false)
	assert.Error(t, err)
}

func TestServeMux(t *testing.T) {
	// setup
	app := mockApp(t)

	// success
	smux := app.ServeMux()
	assert.NotNil(t, smux)
}

func TestServer(t *testing.T) {
	// setup
	app := mockApp(t)

	// success
	serv := app.Server()
	assert.NotNil(t, serv)
}
