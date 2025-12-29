// Package bolt implements database handling functions.
package bolt

import (
	"time"

	"go.etcd.io/bbolt"
)

// Connect returns a new database connection.
func Connect(path string) (*bbolt.DB, error) {
	return bbolt.Open(path, 0644, &bbolt.Options{
		Timeout: 10 * time.Second,
	})
}

// Delete deletes an existing database entry.
func Delete(db *bbolt.DB, name string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(name))
	})
}

// Exists returns true if a database entry exists.
func Exists(db *bbolt.DB, name string) (bool, error) {
	var okay bool
	return okay, db.View(func(tx *bbolt.Tx) error {
		okay = tx.Bucket([]byte(name)) != nil
		return nil
	})
}

// Get returns a value from an existing database entry.
func Get(db *bbolt.DB, name, attr string) (string, error) {
	var data string
	return data, db.View(func(tx *bbolt.Tx) error {
		if buck := tx.Bucket([]byte(name)); buck != nil {
			data = string(buck.Get([]byte(attr)))
		}

		return nil
	})
}

// Set sets a new or existing database entry from a map.
func Set(db *bbolt.DB, name string, pairs map[string]string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}

		for attr, data := range pairs {
			if err := buck.Put([]byte(attr), []byte(data)); err != nil {
				return err
			}
		}

		return nil
	})
}
