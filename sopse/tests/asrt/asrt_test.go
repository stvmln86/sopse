package asrt

import (
	"errors"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	// setup
	err := errors.New("error")

	// success
	okay := Error(t, err, "%s", "error")
	assert.True(t, okay)
}

func TestRow(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success - integer
	okay := Row(t, db, "select 123", 123)
	assert.True(t, okay)

	// success - default
	okay = Row(t, db, "select 'abc'", "abc")
	assert.True(t, okay)
}
