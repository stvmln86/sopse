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

// Get returns a database value.
func Get(db *bbolt.DB, name, attr string) string {
	var data string
	db.View(func(tx *bbolt.Tx) error {
		data = string(tx.Bucket([]byte(name)).Get([]byte(attr)))
		return nil
	})

	return data
}

// MockDB returns a temporary database containing mockData.
func MockDB(t *testing.T) *bbolt.DB {
	dire := t.TempDir()
	dest := filepath.Join(dire, "mock.db")
	db, _ := bbolt.Open(dest, 0644, nil)
	db.Update(func(tx *bbolt.Tx) error {
		for name, pairs := range mockData {
			buck, _ := tx.CreateBucketIfNotExists([]byte(name))
			for attr, data := range pairs {
				buck.Put([]byte(attr), []byte(data))
			}
		}

		return nil
	})

	return db
}
