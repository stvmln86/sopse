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
	var okay bool
	err = db.Get(&okay, "pragma foreign_keys")
	assert.True(t, okay)
	assert.NoError(t, err)
}

func TestSchema(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Schema)
	assert.NoError(t, err)

	// confirm - database
	var size int
	err = db.Get(&size, "select count(*) from SQLITE_SCHEMA")
	assert.NotZero(t, size)
	assert.NoError(t, err)
}

func TestConnect(t *testing.T) {
	// setup
	fileErr := `unable to open database file: no such file or directory`
	textErr := `near "nope": syntax error`

	// success
	db, err := Connect(":memory:", "create table Mock (a)")
	assert.NotNil(t, db)
	assert.NoError(t, err)

	// confirm - database
	var size int
	err = db.Get(&size, "select count(*) from SQLITE_SCHEMA")
	assert.NotZero(t, size)
	assert.NoError(t, err)

	// error - file does not exist
	db, err = Connect("/nope.db", "")
	assert.Nil(t, db)
	test.AssertError(t, err, `cannot connect database "/nope.db" - %s`, fileErr)

	// error - bad schema
	db, err = Connect(":memory:", "nope")
	assert.Nil(t, db)
	test.AssertError(t, err, `cannot connect database ":memory:" - %s`, textErr)
}
