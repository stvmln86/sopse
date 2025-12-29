// Package pair implements the Pair type and methods.
package pair

import (
	"fmt"
	"time"

	"github.com/stvmln86/sopse/sopse/tools/bolt"
	"github.com/stvmln86/sopse/sopse/tools/neat"
	"go.etcd.io/bbolt"
)

// Pair is a single recorded key-value pair.
type Pair struct {
	DB   *bbolt.DB
	Path string
	Body string
	Init time.Time
}

// New returns a new Pair.
func New(db *bbolt.DB, path, body string, init time.Time) *Pair {
	return &Pair{db, path, body, init}
}

// Get returns an existing Pair or nil.
func Get(db *bbolt.DB, uuid, name string) (*Pair, error) {
	path := bolt.Join("pair", uuid, name)
	bmap, err := bolt.Get(db, path)
	switch {
	case bmap == nil:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot get Pair %q - %w", path, err)
	}

	init := neat.Time(bmap["init"])
	return New(db, path, bmap["body"], init), nil
}

// Set sets and returns a new or existing Pair.
func Set(db *bbolt.DB, uuid, name, body string) (*Pair, error) {
	path := bolt.Join("pair", uuid, name)
	pair := New(db, path, body, time.Now())
	if err := bolt.Set(db, path, pair.Map()); err != nil {
		return nil, fmt.Errorf("cannot set Pair %q - %w", path, err)
	}

	return pair, nil
}

// Delete deletes the Pair.
func (p *Pair) Delete() error {
	if err := bolt.Delete(p.DB, p.Path); err != nil {
		return fmt.Errorf("cannot delete Pair %q - %w", p.Path, err)
	}

	return nil
}

// Expired returns true if the Pair is past a duration.
func (p *Pair) Expired(duration time.Duration) bool {
	return neat.Expired(p.Init, duration)
}

// Map returns the Pair as a string map.
func (p *Pair) Map() map[string]string {
	return map[string]string{
		"body": p.Body,
		"init": neat.Unix(p.Init),
	}
}

// Name returns the Pair's path name.
func (p *Pair) Name() string {
	_, _, name := bolt.Split(p.Path)
	return name
}

// Update overwrites the Pair with a new body.
func (p *Pair) Update(body string) error {
	p.Body = body
	if err := bolt.Set(p.DB, p.Path, p.Map()); err != nil {
		return fmt.Errorf("cannot update Pair %q - %w", p.Path, err)
	}

	return nil
}
