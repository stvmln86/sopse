// Package pair implements the Pair type and methods.
package pair

import (
	"fmt"
	"strings"
	"time"

	"github.com/stvmln86/sopse/sopse/tools/bolt"
	"github.com/stvmln86/sopse/sopse/tools/neat"
	"go.etcd.io/bbolt"
)

// Pair is a single recorded key-value pair.
type Pair struct {
	DB   *bbolt.DB
	Addr string
	Body string
	Init time.Time
}

// New returns a new Pair.
func New(db *bbolt.DB, addr, body string, init time.Time) *Pair {
	return &Pair{db, addr, body, init}
}

// Get returns an existing Pair or nil.
func Get(db *bbolt.DB, addr string) (*Pair, error) {
	bmap, err := bolt.Get(db, addr)
	switch {
	case bmap == nil:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot get Pair %q - %w", addr, err)
	}

	init := neat.Time(bmap["init"])
	return New(db, addr, bmap["body"], init), nil
}

// Set sets and returns a new or existing Pair.
func Set(db *bbolt.DB, addr, body string, init time.Time) (*Pair, error) {
	pair := New(db, addr, body, init)
	if err := bolt.Set(db, addr, pair.Map()); err != nil {
		return nil, fmt.Errorf("cannot set Pair %q - %w", addr, err)
	}

	return pair, nil
}

// Delete deletes the Pair.
func (p *Pair) Delete() error {
	if err := bolt.Delete(p.DB, p.Addr); err != nil {
		return fmt.Errorf("cannot delete Pair %q - %w", p.Addr, err)
	}

	return nil
}

// Map returns the Pair as a string map.
func (p *Pair) Map() map[string]string {
	return map[string]string{
		"body": p.Body,
		"init": neat.Unix(p.Init),
	}
}

// Name returns the Pair's base name.
func (p *Pair) Name() string {
	elems := strings.Split(p.Addr, ".")
	return elems[len(elems)-1]
}

// Update overwrites the Pair with a new body.
func (p *Pair) Update(body string) error {
	p.Body = body
	if err := bolt.Set(p.DB, p.Addr, p.Map()); err != nil {
		return fmt.Errorf("cannot update Pair %q - %w", p.Addr, err)
	}

	return nil
}
