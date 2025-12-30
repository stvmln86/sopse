package dbse

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func TestPragma(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Pragma)
	assert.NoError(t, err)

	// confirm - database
	okay := test.Get(t, db, "pragma foreign_keys")
	assert.NotZero(t, okay)
}

func TestSchema(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Schema)
	assert.NoError(t, err)

	// confirm - database
	size := test.Get(t, db, "select count(*) from SQLITE_SCHEMA")
	assert.NotZero(t, size)
}

func TestConnect(t *testing.T) {
	// success
	db, err := Connect(":memory:")
	assert.NotNil(t, db)
	assert.NoError(t, err)
}
