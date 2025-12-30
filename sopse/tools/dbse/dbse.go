// Package dbse implements database constants and functions.
package dbse

import (
	"fmt"

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
	create table if not exists Notes (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		flag integer not null default 0,
		name text    not null,

		check (flag in (0, 1, 2, 3)),
		unique(name)
	);

	create table if not exists Pages (
		id   integer primary key asc,
		init integer not null default (unixepoch()),
		note integer not null,
		body text    not null,
		hash text    not null,

		foreign key (note) references Notes(id) on delete cascade,
		unique(note, hash)
	);

	create index if not exists NoteNames on Notes(name);
`

// Connect returns a new database connection with an executed query.
func Connect(path, code string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("cannot connect database %q - %w", path, err)
	}

	if _, err := db.Exec(code); err != nil {
		return nil, fmt.Errorf("cannot connect database %q - %w", path, err)
	}

	return db, nil
}
