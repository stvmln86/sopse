// Package dbse implements SQLite database schema and functions.
package dbse

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Pragma is the default always-enabled database pragma.
const Pragma = `
	pragma encoding = 'utf-8';
	pragma foreign_keys = true;
`

// Schema is the default first-run database schema.
const Schema = `
	create table if not exists Users (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		uuid text    not null default (lower(hex(randomblob(16)))),
		addr text    not null,

		unique(uuid)
	);

	create table if not exists Pairs (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		user integer not null,
		name text    not null,
		body text    not null,

		foreign key (user) references Users(id) on delete cascade,
		unique(user, name)
	);

	create index if not exists UserUUIDs on Users(uuid);
	create index if not exists PairNames on Pairs(user, name);
`

// Connect returns a new database connection.
func Connect(path string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(Pragma + Schema); err != nil {
		return nil, err
	}

	return db, nil
}
