// Package user implements the User type and methods.
package user

import (
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/stvmln86/sopse/sopse/items/pair"
	"github.com/stvmln86/sopse/sopse/tools/neat"
)

// User is a single recorded user.
type User struct {
	DB   *sqlx.DB `db:"-"`
	ID   int64    `db:"id"`
	Init int64    `db:"init"`
	UUID string   `db:"uuid"`
	Addr string   `db:"addr"`
}

const (
	deleteUser = `
		delete from Users where id=?
	`

	insertUser = `
		insert into Users (addr) values (?) returning *
	`

	selectPairs = `
		select name from Pairs join Users on Pairs.user = Users.id
		where Users.id=? order by name asc
	`

	selectUser = `
		select * from Users where uuid=? limit 1
	`
)

// Create creates and returns a new User.
func Create(db *sqlx.DB, r *http.Request) (*User, error) {
	addr := neat.Addr(r)
	user := &User{DB: db}
	if err := db.Get(user, insertUser, addr); err != nil {
		return nil, err
	}

	return user, nil
}

// Get returns an existing User, or nil.
func Get(db *sqlx.DB, uuid string) (*User, error) {
	user := &User{DB: db}
	err := db.Get(user, selectUser, uuid)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return user, nil
	}
}

// Delete deletes the User.
func (u *User) Delete() error {
	_, err := u.DB.Exec(deleteUser, u.ID)
	return err
}

// GetPair returns an existing Pair from the User.
func (u *User) GetPair(name string) (*pair.Pair, error) {
	return pair.Get(u.DB, u.ID, name)
}

// ListPairs returns all the User's existing Pair names.
func (u *User) ListPairs() ([]string, error) {
	var names []string
	err := u.DB.Select(&names, selectPairs, u.ID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return names, nil
}

// SetPair sets and returns a new or existing Pair from the User.
func (u *User) SetPair(name, body string) (*pair.Pair, error) {
	return pair.Set(u.DB, u.ID, name, body)
}
