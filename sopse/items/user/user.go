// Package user implements the User type and methods.
package user

import (
	"sync"

	"github.com/stvmln86/sopse/sopse/items/pair"
)

// User is a single stored user file.
type User struct {
	Orig  string       `json:"-"`
	mutex sync.RWMutex `json:"-"`

	Addr  string                `json:"addr"`
	Hash  string                `json:"hash"`
	Init  int64                 `json:"init"`
	Pairs map[string]*pair.Pair `json:"pairs"`
}
