package pair

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
	"go.etcd.io/bbolt"
)

var mockTime = time.Unix(1000, 0).Local()

func mockPair(t *testing.T) *Pair {
	db := test.DB(t)
	pair, err := Get(db, "mockUser1", "alpha")
	test.Try(t, err)
	return pair
}

func TestNew(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	pair := New(db, "addr", "body", mockTime)
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, "addr", pair.Addr)
	assert.Equal(t, "body", pair.Body)
	assert.Equal(t, mockTime, pair.Init)
}

func TestGet(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	pair, err := Get(db, "mockUser1", "alpha")
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, "pair.mockUser1.alpha", pair.Addr)
	assert.Equal(t, "Alpha.", pair.Body)
	assert.Equal(t, mockTime, pair.Init)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	pair, err := Set(db, "mockUser1", "alpha", "body", mockTime)
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, "pair.mockUser1.alpha", pair.Addr)
	assert.Equal(t, "body", pair.Body)
	assert.Equal(t, mockTime, pair.Init)
	assert.NoError(t, err)

	// confirm - database
	test.Try(t, db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("pair.mockUser1.alpha"))
		for attr, want := range pair.Map() {
			data := string(buck.Get([]byte(attr)))
			assert.Equal(t, want, data)
		}

		return nil
	}))
}

func TestDelete(t *testing.T) {
	// setup
	pair := mockPair(t)

	// success
	err := pair.Delete()
	assert.NoError(t, err)

	// confirm - database
	test.Try(t, pair.DB.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("pair.mockUser1.alpha"))
		assert.Nil(t, buck)
		return nil
	}))
}

func TestExpired(t *testing.T) {
	// setup
	pair := mockPair(t)

	// success
	okay := pair.Expired(1 * time.Second)
	assert.True(t, okay)
}

func TestMap(t *testing.T) {
	// setup
	pair := mockPair(t)

	// success
	bmap := pair.Map()
	assert.Equal(t, map[string]string{
		"body": "Alpha.",
		"init": "1000",
	}, bmap)
}

func TestName(t *testing.T) {
	// setup
	pair := mockPair(t)

	// success
	name := pair.Name()
	assert.Equal(t, "alpha", name)
}

func TestUpdate(t *testing.T) {
	// setup
	pair := mockPair(t)

	// success
	err := pair.Update("body")
	assert.NoError(t, err)

	// confirm - database
	test.Try(t, pair.DB.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("pair.mockUser1.alpha"))
		body := string(buck.Get([]byte("body")))
		assert.Equal(t, "body", body)
		return nil
	}))
}
