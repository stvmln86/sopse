// Package user implements the User type and methods.
package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/stvmln86/sopse/sopse/items/pair"
	"github.com/stvmln86/sopse/sopse/tools/bolt"
	"github.com/stvmln86/sopse/sopse/tools/neat"
	"go.etcd.io/bbolt"
)

type User struct {
	DB   *bbolt.DB
	Addr string
	From string
	Init time.Time
}

// New returns a new User.
func New(db *bbolt.DB, addr, from string, init time.Time) *User {
	return &User{db, addr, from, init}
}

// Create creates and returns a new
func Create(db *bbolt.DB, r *http.Request) (*User, error) {
	addr := bolt.Join("user", neat.UUID())
	from := neat.From(r)
	user := New(db, addr, from, time.Now())
	if err := bolt.Set(db, addr, user.Map()); err != nil {
		return nil, fmt.Errorf("cannot create User %q - %w", addr, err)
	}

	return user, nil
}

// Get returns an existing User or nil.
func Get(db *bbolt.DB, uuid string) (*User, error) {
	addr := bolt.Join("user", uuid)
	bmap, err := bolt.Get(db, addr)
	switch {
	case bmap == nil:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot get User %q - %w", addr, err)
	}

	init := neat.Time(bmap["init"])
	return New(db, addr, bmap["from"], init), nil
}

// Delete deletes the User.
func (u *User) Delete() error {
	if err := bolt.Delete(u.DB, u.Addr); err != nil {
		return fmt.Errorf("cannot delete User %q - %w", u.Addr, err)
	}

	return nil
}

// GetPair returns an existing Pair from the User or nil
func (u *User) GetPair(name string) (*pair.Pair, error) {
	return pair.Get(u.DB, u.UUID(), name)
}

// ListPairs returns the addresses of all the User's existing Pairs.
func (u *User) ListPairs() ([]string, error) {
	pref := bolt.Join("pair", u.UUID())
	addrs, err := bolt.List(u.DB, pref)
	if err != nil {
		return nil, fmt.Errorf("cannot list User %q - %w", u.Addr, err)
	}

	return addrs, nil
}

// Map returns the User as a map.
func (u *User) Map() map[string]string {
	return map[string]string{
		"from": u.From,
		"init": neat.Unix(u.Init),
	}
}

// SetPair sets and returns a new or existing Pair for the User.
func (u *User) SetPair(name, body string) (*pair.Pair, error) {
	return pair.Set(u.DB, u.UUID(), name, body)
}

// UUID returns the User's base name UUID.
func (u *User) UUID() string {
	elems := strings.Split(u.Addr, ".")
	return elems[len(elems)-1]
}
