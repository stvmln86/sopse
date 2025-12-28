// Package book implements the Book type and methods.
package book

import (
	"sync"
)

// Book is a directory of stored user files.
type Book struct {
	Dire  string        `json:"-"`
	mutex *sync.RWMutex `json:"-"`
}
