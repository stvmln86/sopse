// Package user implements the User type and methods.
package user

import (
	"errors"
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
	Path string
	Addr string
	Init time.Time
}

// New returns a new User.
func New(db *bbolt.DB, path, addr string, init time.Time) *User {
	return &User{db, path, addr, init}
}

// Create creates and returns a new User.
func Create(db *bbolt.DB, r *http.Request) (*User, error) {
	path := bolt.Join("user", neat.UUID())
	addr := neat.Addr(r)
	user := New(db, path, addr, time.Now())
	if err := bolt.Set(db, path, user.Map()); err != nil {
		return nil, fmt.Errorf("cannot create User %q - %w", path, err)
	}

	return user, nil
}

// Get returns an existing User or nil.
func Get(db *bbolt.DB, uuid string) (*User, error) {
	path := bolt.Join("user", uuid)
	bmap, err := bolt.Get(db, path)
	switch {
	case bmap == nil:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot get User %q - %w", path, err)
	}

	init := neat.Time(bmap["init"])
	return New(db, path, bmap["addr"], init), nil
}

// Delete deletes the User.
func (u *User) Delete() error {
	if err := bolt.Delete(u.DB, u.Path); err != nil {
		return fmt.Errorf("cannot delete User %q - %w", u.Path, err)
	}

	return nil
}

// GetPair returns an existing Pair from the User or nil.
func (u *User) GetPair(name string) (*pair.Pair, error) {
	return pair.Get(u.DB, u.UUID(), name)
}

// ListPairs returns the User's existing Pairs.
func (u *User) ListPairs() ([]*pair.Pair, error) {
	pref := bolt.Join("pair", u.UUID())
	paths, err := bolt.List(u.DB, pref)
	if err != nil {
		return nil, fmt.Errorf("cannot list User %q - %w", u.Path, err)
	}

	var pairs []*pair.Pair
	for _, path := range paths {
		_, uuid, name := bolt.Split(path)
		pair, err := pair.Get(u.DB, uuid, name)
		if err != nil {
			sub := errors.Unwrap(err)
			return nil, fmt.Errorf("cannot list User %q - %w", u.Path, sub)
		}

		pairs = append(pairs, pair)
	}

	return pairs, nil
}

// Map returns the User as a map.
func (u *User) Map() map[string]string {
	return map[string]string{
		"addr": u.Addr,
		"init": neat.Unix(u.Init),
	}
}

// SetPair sets and returns a new or existing Pair into the User.
func (u *User) SetPair(name, body string) (*pair.Pair, error) {
	return pair.Set(u.DB, u.UUID(), name, body)
}

// UUID returns the User's path UUID.
func (u *User) UUID() string {
	elems := strings.Split(u.Path, ".")
	return elems[len(elems)-1]
}
