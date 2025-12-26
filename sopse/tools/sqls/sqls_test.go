package sqls

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
)

func TestPragma(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Pragma)
	assert.NoError(t, err)

	// confirm - pragma
	fkey := asrt.Get(t, db, "pragma foreign_keys")
	assert.Equal(t, int64(1), fkey)
}

func TestSchema(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Schema)
	assert.NoError(t, err)

	// confirm - schema
	size := asrt.Get(t, db, "select count(*) from SQLITE_SCHEMA")
	assert.NotZero(t, size)
}
