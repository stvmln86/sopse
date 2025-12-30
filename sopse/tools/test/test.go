// Package test implements unit testing data and functions.
package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// Get returns the first value from a database query.
func Get(t *testing.T, db *sqlx.DB, code string, elems ...any) any {
	var data any
	err := db.Get(&data, code, elems...)
	assert.NoError(t, err)
	return data
}
