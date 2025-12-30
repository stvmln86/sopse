// Package test implements unit testing data and functions.
package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/dbse"
)

// mockData is additional database data for unit testing.
const mockData = `
	insert into Users (init, uuid, addr) values
		(1000, 'mockUser', '1.1.1.1');

	insert into Pairs (init, user, name, body) values
		(1000, 1, 'alpha', 'Alpha.'),
		(1000, 1, 'bravo', 'Bravo.');
`

// DB returns a new in-memory database populated with mock data.
func DB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec(dbse.Pragma + dbse.Schema + mockData)
	return db
}

// Get returns the first value from a database query.
func Get(t *testing.T, db *sqlx.DB, code string, elems ...any) any {
	var data any
	err := db.Get(&data, code, elems...)
	assert.NoError(t, err)
	return data
}
