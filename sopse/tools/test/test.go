package test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stvmln86/sopse/sopse/tools/sqls"
)

// MockData is additional database data for unit testing.
const MockData = `
	insert into Users (init, uuid, addr) values (1000, '1111', '1.1.1.1');
	insert into Pairs (init, user, name, body) values
		(1000, 1, 'alpha', 'Alpha.' || char(10)),
		(2000, 1, 'bravo', 'Bravo.' || char(10));
`

// Get returns the first value from a database query.
func Get(t *testing.T, db *sqlx.DB, code string, elems ...any) any {
	var data any
	if err := db.Get(&data, code, elems...); err != nil {
		t.Fatal(err)
	}

	return data
}

// MockDB returns an in-memory database populated with MockData.
func MockDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	if _, err := db.Exec(sqls.Pragma + sqls.Schema + MockData); err != nil {
		t.Fatal(err)
	}

	return db
}
