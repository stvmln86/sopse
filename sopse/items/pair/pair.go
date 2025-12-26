// Package pair implements the Pair type and methods.
package pair

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Pair is a single recorded pair in a database.
type Pair struct {
	DB   *sqlx.DB `db:"-"`
	ID   int64    `db:"id"`
	Init int64    `db:"init"`
	User int64    `db:"user"`
	Name string   `db:"name"`
	Body string   `db:"body"`
}

const (
	deletePair       = "delete from Pairs where id=?"
	insertPair       = "insert into Pairs (user, name, body) values (?, ?, ?)"
	selectPairExists = "select exists(select 1 from Pairs where id=?)"
	selectPairID     = "select * from Pairs where id=? limit 1"
	selectPairName   = "select * from Pairs where user=? and name=? limit 1"
)

// Create creates and returns a new Pair.
func Create(db *sqlx.DB, user int64, name, body string) (*Pair, error) {
	rslt, err := db.Exec(insertPair, user, name, body)
	if err != nil {
		return nil, fmt.Errorf("cannot create Pair %q - %w", name, err)
	}

	last, err := rslt.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("cannot create Pair %q - %w", name, err)
	}

	pair := &Pair{DB: db}
	if err := db.Get(pair, selectPairID, last); err != nil {
		return nil, fmt.Errorf("cannot create Pair %q - %w", name, err)
	}

	return pair, nil
}

// Get returns an existing Pair by user and name.
func Get(db *sqlx.DB, user int64, name string) (*Pair, error) {
	pair := &Pair{DB: db}
	err := db.Get(pair, selectPairName, user, name)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot get Pair %q - %w", name, err)
	default:
		return pair, nil
	}
}

// Delete deletes the Pair.
func (p *Pair) Delete() error {
	if _, err := p.DB.Exec(deletePair, p.ID); err != nil {
		return fmt.Errorf("cannot delete Pair %q - %w", p.Name, err)
	}

	return nil
}

// Exists returns true if the Pair exists.
func (p *Pair) Exists() (bool, error) {
	var ok bool
	if err := p.DB.Get(&ok, selectPairExists, p.ID); err != nil {
		return false, fmt.Errorf("cannot read Pair %q - %w", p.Name, err)
	}

	return ok, nil
}
