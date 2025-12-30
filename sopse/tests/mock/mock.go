// Package mock implements unit testing mock data and functions.
package mock

import (
	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/sopse/sopse/tools/dbse"
)

// Inserts is mock database data for unit testing.
const Inserts = `
	insert into Notes (init, flag, name) values
		(unixepoch('2026-01-01 12:00:00'), 0, 'alpha'),
		(unixepoch('2026-01-02 12:00:00'), 1, 'bravo'),
		(unixepoch('2026-01-03 12:00:00'), 2, 'charlie'),
		(unixepoch('2026-01-04 12:00:00'), 3, 'delta');

	insert into Pages (init, note, body, hash) values
		(unixepoch('2026-01-01 12:00:00'), 1, 'Alpha old.' || char(10), '988nowOiHm-63RWHAZHliwsSg_aiQ0gw9PQ0AqLRRG0'),
		(unixepoch('2026-01-01 13:00:00'), 1, 'Alpha new.' || char(10), 'oE39AcPiuCRhVTo_8oY1KXHsoA12gx1l-Cnvm_-REPY'),
		(unixepoch('2026-01-02 12:00:00'), 2, 'Bravo.'     || char(10), 'hBGHHe7shp2EUSw0FNyjNYmIHrrJWohrxL1aThcr_nw'),
		(unixepoch('2026-01-03 12:00:00'), 3, 'Charlie.'   || char(10), 'D-i5pKwHJlJzJOSnhDE76_leokWs6UxRaHqK5_uGcaE'),
		(unixepoch('2026-01-04 12:00:00'), 4, 'Delta.'     || char(10), '___________________________________________');
`

// DB returns an in-memory database with pragma, schema and mock data.
func DB() *sqlx.DB {
	db := sqlx.MustConnect("sqlite3", ":memory:")
	db.MustExec(dbse.Pragma + dbse.Schema + Inserts)
	return db
}
