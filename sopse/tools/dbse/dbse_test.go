package dbse

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	// success
	db, err := Connect(":memory:")
	assert.NotNil(t, db)
	assert.NoError(t, err)
}

func TestPragma(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Pragma)
	assert.NoError(t, err)
}

func TestSchema(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	_, err := db.Exec(Schema)
	assert.NoError(t, err)
}
