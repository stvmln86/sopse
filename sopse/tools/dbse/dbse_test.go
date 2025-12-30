package dbse

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

	// confirm - database
	asrt.Row(t, db, "pragma foreign_keys", 1)
}

func TestSchema(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Schema)
	assert.NoError(t, err)

	// confirm - database
	asrt.Row(t, db, "select count(*) from SQLITE_SCHEMA", 5)
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
	asrt.Row(t, db, "select count(*) from SQLITE_SCHEMA", 1)

	// error - file does not exist
	db, err = Connect("/nope.db", "")
	assert.Nil(t, db)
	asrt.Error(t, err, `cannot connect database "/nope.db" - %s`, fileErr)

	// error - bad schema
	db, err = Connect(":memory:", "nope")
	assert.Nil(t, db)
	asrt.Error(t, err, `cannot connect database ":memory:" - %s`, textErr)
}
