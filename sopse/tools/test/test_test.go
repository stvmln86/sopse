package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	// success
	db := DB()
	assert.NotNil(t, db)

	// confirm - database
	var size int
	err := db.Get(&size, "select count(*) from Users")
	assert.NotZero(t, size)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	data := Get(t, db, "select 123")
	assert.Equal(t, int64(123), data)
}
