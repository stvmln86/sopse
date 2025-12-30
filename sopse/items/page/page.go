// Package page implements the Page type and methods.
package page

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/sopse/sopse/tools/neat"
)

// Page is a single historical version of a recorded note.
type Page struct {
	db   *sqlx.DB
	ID   int64  `db:"id"`
	Init int64  `db:"init"`
	Note int64  `db:"note"`
	Body string `db:"body"`
	Hash string `db:"hash"`
}

const (
	create = `insert into Pages (note, body, hash) values (?, ?, ?) returning *`
	latest = `select * from Pages where note=? order by id desc limit 1`
)

// Create creates and returns a new Page for a Note.
func Create(db *sqlx.DB, note int64, body string) (*Page, error) {
	body = neat.Body(body)
	hash := neat.Hash(body)
	page := &Page{db: db}
	if err := db.Get(page, create, note, body, hash); err != nil {
		return nil, fmt.Errorf("cannot create Page - %w", err)
	}

	return page, nil
}

// Latest returns the latest Page for an existing Note.
func Latest(db *sqlx.DB, note int64) (*Page, error) {
	page := &Page{db: db}
	err := db.Get(page, latest, note)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot get Page - %w", err)
	default:
		return page, nil
	}
}

// InitTime returns the Page's creation time.
func (p *Page) InitTime() time.Time {
	return neat.Time(p.Init)
}

// Verify returns true if the Page's body matches its hash.
func (p *Page) Verify() bool {
	return neat.Hash(p.Body) == p.Hash
}
