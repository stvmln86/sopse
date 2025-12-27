// Package pair implements the Pair type and methods.
package pair

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Pair is a single recorded key-value pair in a database.
type Pair struct {
	DB   *sqlx.DB `db:""`
	ID   int64    `db:"id"`
	Body string   `db:"body"`
	Init int64    `db:"init"`
	Name string   `db:"name"`
	User int64    `db:"user"`
}

const (
	delete       = "delete from Pairs where id=?"
	selectExists = "select exists (select 1 from Pairs where id=?)"
	selectName   = "select * from Pairs where user=? and name=? limit 1"
	selectID     = "select * from Pairs where id=? limit 1"
	upsert       = `
		insert into Pairs (user, name, body) values (?, ?, ?)
		on conflict(user, name) do update set body=excluded.body
	`
)

// Get returns an existing Pair by name.
func Get(db *sqlx.DB, user int64, name string) (*Pair, error) {
	pair := &Pair{DB: db}
	err := db.Get(pair, selectName, user, name)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot read Pair %q - %w", name, err)
	default:
		return pair, nil
	}
}

// Set creates a new Pair or overwrites an existing Pair.
func Set(db *sqlx.DB, user int64, name, body string) (*Pair, error) {
	if _, err := db.Exec(upsert, user, name, body); err != nil {
		return nil, fmt.Errorf("cannot set Pair %q - %w", name, err)
	}

	pair := &Pair{DB: db}
	if err := db.Get(pair, selectName, user, name); err != nil {
		return nil, fmt.Errorf("cannot set Pair %q - %w", name, err)
	}

	return pair, nil
}

// Delete deletes the Pair.
func (p *Pair) Delete() error {
	if _, err := p.DB.Exec(delete, p.ID); err != nil {
		return fmt.Errorf("cannot delete Pair %q - %w", p.Name, err)
	}

	return nil
}

// Exists returns true if the Pair exists.
func (p *Pair) Exists() (bool, error) {
	var okay bool
	if err := p.DB.Get(&okay, selectExists, p.ID); err != nil {
		return false, fmt.Errorf("cannot read Pair %q - %w", p.Name, err)
	}

	return okay, nil
}
