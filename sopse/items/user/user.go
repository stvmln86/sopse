// Package user implements the User type and methods.
package user

import (
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/stvmln86/sopse/sopse/items/pair"
	"github.com/stvmln86/sopse/sopse/tools/neat"
)

// User is a single stored user file.
type User struct {
	Orig  string        `json:"-"`
	mutex *sync.RWMutex `json:"-"`

	Addr  string                `json:"addr"`
	Hash  string                `json:"hash"`
	Init  int64                 `json:"init"`
	Pairs map[string]*pair.Pair `json:"pairs"`
}

// Create creates and returns a new User.
func Create(dire, addr, salt string) (*User, error) {
	init := time.Now().Unix()
	unix := strconv.FormatInt(init, 10)
	hash := neat.Hash(addr, unix, salt)
	user := &User{
		Orig:  filepath.Join(dire, hash+".json"),
		mutex: new(sync.RWMutex),
		Addr:  addr,
		Hash:  hash,
		Init:  init,
		Pairs: make(map[string]*pair.Pair),
	}

	if err := file.Create(user.Orig, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Read returns an existing user.
func Read(dire, hash string) (*User, error) {
	orig := filepath.Join(dire, hash+".json")
	user := &User{Orig: orig}
	if err := file.Read(orig, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Delete deletes the User.
func (u *User) Delete() error {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	return file.Delete(u.Orig)
}

// DeletePair deletes a Pair from the User.
func (u *User) DeletePair(name string) error {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	delete(u.Pairs, name)
	return file.Update(u.Orig, u)
}

// Exists returns true if the User exists.
func (u *User) Exists() bool {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	return file.Exists(u.Orig)
}

// GetPair returns an existing Pair from the User.
func (u *User) GetPair(name string) (*pair.Pair, bool) {
	u.mutex.RLock()
	defer u.mutex.RUnlock()
	pair, okay := u.Pairs[name]
	return pair, okay
}

// SetPair sets a new or existing Pair into the User.
func (u *User) SetPair(name, body string) (*pair.Pair, error) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.Pairs[name] = pair.New(body)
	return u.Pairs[name], file.Update(u.Orig, u)
}

// Update updates the User's file.
func (u *User) Update() error {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	return file.Update(u.Orig, u)
}
