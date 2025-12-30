// Package note implements the Note type and methods.
package note

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/sopse/sopse/items/page"
	"github.com/stvmln86/sopse/sopse/tools/neat"
)

// Note is a metadata root for a recorded note.
type Note struct {
	db   *sqlx.DB
	ID   int64  `db:"id"`
	Init int64  `db:"init"`
	Flag int64  `db:"flag"`
	Name string `db:"name"`
}

// OKAY indicates a Note has no status.
const OKAY int64 = 0

// LOCK indicates a Note cannot be overwritten.
const LOCK int64 = 1

// SDEL indicates a Note is soft-deleted.
const SDEL int64 = 2

// WARN indicates a Note has an internal issue.
const WARN int64 = 3

const (
	create = `insert into Notes (flag, name) values (?, ?) returning *`
	read   = `select * from Notes where name=? limit 1`
)

// Create creates and returns a new Note.
func Create(db *sqlx.DB, name string) (*Note, error) {
	name = neat.Name(name)
	note := &Note{db: db}
	if err := db.Get(note, create, OKAY, name); err != nil {
		return nil, fmt.Errorf("cannot create Note - %w", err)
	}

	return note, nil
}

// Read returns an existing Page.
func Read(db *sqlx.DB, name string) (*Note, error) {
	name = neat.Name(name)
	note := &Note{db: db}
	if err := db.Get(note, read, name); err != nil {
		return nil, fmt.Errorf("cannot create Note - %w", err)
	}

	return note, nil
}

// InitTime returns the Note's creation time.
func (n *Note) InitTime() time.Time {
	return neat.Time(n.Init)
}

// Latest returns the Note's latest Page.
func (n *Note) Latest() (*page.Page, error) {
	return page.Latest(n.db, n.ID)
}
