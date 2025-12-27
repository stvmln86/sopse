// Package user implements the User type and methods.
package user

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/sopse/sopse/items/pair"
)

// User is a single recorded user in the database.
type User struct {
	DB   *sqlx.DB `db:""`
	ID   int64    `db:"id"`
	Addr string   `db:"addr"`
	Init int64    `db:"init"`
	UUID string   `db:"uuid"`
}

const (
	delete       = "delete from Users where id=?"
	insert       = "insert into Users (addr) values (?)"
	selectExists = "select exists (select 1 from Users where id=?)"
	selectID     = "select * from Users where id=? limit 1"
	selectUUID   = "select * from Users where uuid=? limit 1"
)

// Create creates and returns a new User.
func Create(db *sqlx.DB, addr string) (*User, error) {
	rslt, err := db.Exec(insert, addr)
	if err != nil {
		return nil, fmt.Errorf("cannot create User %q - %w", addr, err)
	}

	last, err := rslt.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("cannot create User %q - %w", addr, err)
	}

	user := &User{DB: db}
	if err := db.Get(user, selectID, last); err != nil {
		return nil, fmt.Errorf("cannot create User %q - %w", addr, err)
	}

	return user, nil
}

// Get returns an existing User by UUID.
func Get(db *sqlx.DB, uuid string) (*User, error) {
	user := &User{DB: db}
	err := db.Get(user, selectUUID, uuid)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, fmt.Errorf("cannot read User %q - %w", uuid, err)
	default:
		return user, nil
	}
}

// Delete deletes the User.
func (u *User) Delete() error {
	if _, err := u.DB.Exec(delete, u.ID); err != nil {
		return fmt.Errorf("cannot delete User %q - %w", u.UUID, err)
	}

	return nil
}

// Exists returns true if the User exists.
func (u *User) Exists() (bool, error) {
	var okay bool
	if err := u.DB.Get(&okay, selectExists, u.ID); err != nil {
		return false, fmt.Errorf("cannot read User %q - %w", u.UUID, err)
	}

	return okay, nil
}

// GetPair returns an existing Pair by name.
func (u *User) GetPair(name string) (*pair.Pair, error) {
	return pair.Get(u.DB, u.ID, name)
}

// SetPair creates a new Pair or overwrites an existing Pair.
func (u *User) SetPair(name, body string) (*pair.Pair, error) {
	return pair.Set(u.DB, u.ID, name, body)
}
