// Package mock implements unit testing mock data and functions.
package mock

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stvmln86/sopse/sopse/tools/sqls"
)

// MockInserts is mock database insert data.
const MockInserts = `
	insert into Users (uuid, init, addr) values
		('aaaa', 1000, '1.1.1.1'),
		('bbbb', 2000, '2.2.2.2');

	insert into Pairs (init, user, name, body) values
		(1000, 1, 'alpha', 'Alpha body.' || char(10)),
		(1100, 1, 'bravo', 'Bravo body.' || char(10));
`

// DB returns an in-memory database populated with MockInserts.
func DB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec(sqls.Pragma + sqls.Schema + MockInserts)
	return db
}

// Request returns a new mock Request.
func Request(meth, path, body string) *http.Request {
	buff := bytes.NewBufferString(body)
	return httptest.NewRequest(meth, path, buff)
}
