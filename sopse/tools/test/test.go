// Package test implements unit testing data and functions.
package test

import (
	"path/filepath"
	"testing"

	"go.etcd.io/bbolt"
)

// MockData is a map of mock database data for unit testing.
var MockData = map[string]map[string]string{
	"user.mockUser1":       {"hash": "mockUser", "addr": "1.1.1.1", "init": "1000"},
	"pair.mockUser1.alpha": {"name": "alpha", "body": "Alpha.", "init": "1000"},
	"pair.mockUser1.bravo": {"name": "bravo", "body": "Bravo.", "init": "1100"},
}

// DB returns a temporary database containing mockData.
func DB(t *testing.T) *bbolt.DB {
	dire := t.TempDir()
	dest := filepath.Join(dire, t.Name()+".db")
	db, err := bbolt.Open(dest, 0644, nil)
	Try(t, err)

	Try(t, db.Update(func(tx *bbolt.Tx) error {
		for name, bmap := range MockData {
			buck, err := tx.CreateBucketIfNotExists([]byte(name))
			if err != nil {
				return err
			}

			for attr, data := range bmap {
				if err := buck.Put([]byte(attr), []byte(data)); err != nil {
					return err
				}
			}
		}

		return nil
	}))

	return db
}

// Try fails a test on a non-nil error.
func Try(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
