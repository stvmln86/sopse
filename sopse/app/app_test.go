package app

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func mockApp(t *testing.T) *App {
	db := test.MockDB(t)
	return New(db)
}

func TestNew(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	app := New(db)
	assert.Equal(t, db, app.DB)
}

func TestNewConnect(t *testing.T) {
	// success
	app, err := NewConnect(":memory:")
	assert.NotNil(t, app.DB)
	assert.NoError(t, err)
}

func TestServer(t *testing.T) {
	// setup
	app := New(nil)

	// success
	serv := app.Server()
	assert.NotEmpty(t, serv.Addr)
	assert.NotNil(t, serv.Handler)
}
