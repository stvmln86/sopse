// Package mock implements unit testing mock data and functions.
package mock

import (
	"path/filepath"
	"testing"

	"go.etcd.io/bbolt"
)

// Data is a map of mock database data for unit testing.
var Data = map[string]map[string]string{
	"user.mockUser1": {
		"addr": "1.1.1.1",
		"init": "1000",
	},

	"pair.mockUser1.alpha": {
		"body": "Alpha.",
		"init": "1000",
	},

	"pair.mockUser1.bravo": {
		"body": "Bravo.",
		"init": "1001",
	},
}

// DB returns a temporary database containing all mock data.
func DB(t *testing.T) *bbolt.DB {
	dire := t.TempDir()
	path := filepath.Join(dire, t.Name()+".db")
	db, _ := bbolt.Open(path, 0600, nil)
	db.Update(func(tx *bbolt.Tx) error {
		for name, bmap := range Data {
			buck, _ := tx.CreateBucketIfNotExists([]byte(name))
			for attr, data := range bmap {
				buck.Put([]byte(attr), []byte(data))
			}
		}

		return nil
	})

	return db
}
