// Package pair implements the Pair type and methods.
package pair

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Pair is a single recorded key-value pair.
type Pair struct {
	DB   *sqlx.DB `db:"-"`
	ID   int64    `db:"id"`
	Init int64    `db:"init"`
	User int64    `db:"user"`
	Name string   `db:"name"`
	Body string   `db:"body"`
}

const (
	deletePair = `
		delete from Pairs where id=?
	`

	insertPair = `
		insert into Pairs (user, name, body) values (?, ?, ?) returning *
	`

	updatePair = `
		update Pairs set body=? where id=?
	`

	selectPair = `
		select * from Pairs where user=? and name=? limit 1
	`
)

// Create creates and returns a new Pair.
func Create(db *sqlx.DB, user int64, name, body string) (*Pair, error) {
	pair := &Pair{DB: db}
	if err := db.Get(pair, insertPair, user, name, body); err != nil {
		return nil, err
	}

	return pair, nil
}

// Get returns an existing Pair, or nil.
func Get(db *sqlx.DB, user int64, name string) (*Pair, error) {
	pair := &Pair{DB: db}
	err := db.Get(pair, selectPair, user, name)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return pair, nil
	}
}

// Delete deletes the Pair.
func (p *Pair) Delete() error {
	_, err := p.DB.Exec(deletePair, p.ID)
	return err
}

// Update overwrites the Pair's body.
func (p *Pair) Update(body string) error {
	if _, err := p.DB.Exec(updatePair, body, p.ID); err != nil {
		return err
	}

	p.Body = body
	return nil
}
