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
func DB(t *testing.T) *bbolt.DB {
	dire := t.TempDir()
	dest := filepath.Join(dire, "mock.db")
	db, err := bbolt.Open(dest, 0644, nil)
	Try(t, err)

	for name, pairs := range mockData {
		for attr, data := range pairs {
			Set(t, db, name, attr, data)
		}
	}

	return db
}

// Get returns a database value.
func Get(t *testing.T, db *bbolt.DB, name, attr string) string {
	var data string
	Try(t, db.View(func(tx *bbolt.Tx) error {
		data = string(tx.Bucket([]byte(name)).Get([]byte(attr)))
		return nil
	}))

	return data
}

// Set sets a new or existing database value.
func Set(t *testing.T, db *bbolt.DB, name, attr, data string) {
	Try(t, db.Update(func(tx *bbolt.Tx) error {
		buck, _ := tx.CreateBucketIfNotExists([]byte(name))
		buck.Put([]byte(attr), []byte(data))
		return nil
	}))
}

// Try fails a test on a non-nil error.
func Try(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
