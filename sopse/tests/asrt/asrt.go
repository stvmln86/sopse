// Package asrt implements unit testing assertion functions.
package asrt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

// Bucket asserts a database bucket is equal to a map.
func Bucket(t *testing.T, db *bbolt.DB, path string, bmap map[string]string) {
	db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte(path))
		assert.NotNil(t, buck)

		for attr, want := range bmap {
			data := string(buck.Get([]byte(attr)))
			assert.Equal(t, want, string(data))
		}

		return nil
	})
}

// NoBucket asserts a database bucket does not exist.
func NoBucket(t *testing.T, db *bbolt.DB, path string) {
	db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte(path))
		assert.Nil(t, buck)
		return nil
	})
}

// TimeNow asserts a Time object is within ten seconds of now.
func TimeNow(t *testing.T, tobj time.Time) {
	assert.WithinDuration(t, time.Now(), tobj, 10*time.Second)
}
