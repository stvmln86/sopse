package mock

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
	"github.com/stvmln86/sopse/sopse/tools/dbse"
)

func TestInserts(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(dbse.Pragma + dbse.Schema + Inserts)
	assert.NoError(t, err)

	// confirm - database
	asrt.Row(t, db, "select count(*) from Notes", 4)
	asrt.Row(t, db, "select count(*) from Pages", 5)
}

func TestDB(t *testing.T) {
	// success
	db := DB()
	assert.NotNil(t, db)
}
