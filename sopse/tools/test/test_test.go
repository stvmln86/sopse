package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// setup
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	// success
	data := Get(t, db, "select ?", 123)
	assert.Equal(t, int64(123), data)
}

func TestMockDB(t *testing.T) {
	// success
	db := MockDB(t)
	assert.NotNil(t, db)
}
