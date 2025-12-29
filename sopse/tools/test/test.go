// Package test implements unit testing data and functions.
package test

import (
	"path/filepath"
	"testing"

	"go.etcd.io/bbolt"
)

// mockData is a map of mock database data for unit testing.
var mockData = map[string]map[string]string{
	"user.mockUser1":       {"addr": "1.1.1.1", "init": "1000"},
	"pair.mockUser1.alpha": {"body": "Alpha.", "init": "1000"},
	"pair.mockUser1.bravo": {"body": "Bravo.", "init": "1100"},
}

// DB returns a temporary database containing mockData.
func DB(t *testing.T, inits ...bool) *bbolt.DB {
	dire := t.TempDir()
	dest := filepath.Join(dire, t.Name()+".db")
	db, err := bbolt.Open(dest, 0644, nil)
	Try(t, err)

	if len(inits) != 0 && inits[0] {
		Try(t, db.Update(func(tx *bbolt.Tx) error {
			for name, pairs := range mockData {
				buck, err := tx.CreateBucketIfNotExists([]byte(name))
				if err != nil {
					return err
				}

				for attr, data := range pairs {
					if err := buck.Put([]byte(attr), []byte(data)); err != nil {
						return err
					}
				}
			}

			return nil
		}))
	}

	return db
}

// Get returns a database value.
func Get(t *testing.T, db *bbolt.DB, name, attr string) string {
	var data string
	Try(t, db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte(name))
		data = string(buck.Get([]byte(attr)))
		return nil
	}))

	return data
}

// Set sets a new or existing database value.
func Set(t *testing.T, db *bbolt.DB, name, attr, data string) {
	Try(t, db.Batch(func(tx *bbolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}

		return buck.Put([]byte(attr), []byte(data))
	}))
}

// Try fails a test on a non-nil error.
func Try(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
