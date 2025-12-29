// Package bolt implements database handling functions.
package bolt

import (
	"strings"
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

// Get returns an existing database entry as a map.
func Get(db *bbolt.DB, name string) (map[string]string, error) {
	var bmap map[string]string
	return bmap, db.View(func(tx *bbolt.Tx) error {
		if buck := tx.Bucket([]byte(name)); buck != nil {
			bmap = make(map[string]string)
			return buck.ForEach(func(attr, data []byte) error {
				bmap[string(attr)] = string(data)
				return nil
			})
		}

		return nil
	})
}

// Join returns a database entry name from one or more elements.
func Join(head string, elems ...string) string {
	elems = append([]string{head}, elems...)
	return strings.Join(elems, ".")
}

// List returns existing database entry names containing a prefix.
func List(db *bbolt.DB, pref string) ([]string, error) {
	var names []string
	return names, db.View(func(tx *bbolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
			if strings.HasPrefix(string(name), pref) {
				names = append(names, string(name))
			}

			return nil
		})
	})
}

// Set sets a new or existing database entry from a map.
func Set(db *bbolt.DB, name string, bmap map[string]string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		buck, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}

		for attr, data := range bmap {
			if err := buck.Put([]byte(attr), []byte(data)); err != nil {
				return err
			}
		}

		return nil
	})
}
