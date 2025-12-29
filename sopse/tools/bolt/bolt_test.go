package bolt

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
	"go.etcd.io/bbolt"
)

func TestConnect(t *testing.T) {
	// setup
	dire := t.TempDir()
	path := filepath.Join(dire, t.Name()+".db")

	// success
	db, err := Connect(path)
	assert.Contains(t, db.String(), "TestConnect.db")
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	err := Delete(db, "user.mockUser1")
	assert.NoError(t, err)

	// confirm - database
	test.Try(t, db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("user.mockUser1"))
		assert.Nil(t, buck)
		return nil
	}))
}

func TestExists(t *testing.T) {
	// setup
	db := test.DB(t)

	// success - true
	okay, err := Exists(db, "user.mockUser1")
	assert.True(t, okay)
	assert.NoError(t, err)

	// success - false
	okay, err = Exists(db, "nope")
	assert.False(t, okay)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.DB(t)

	// success - entry exists
	bmap, err := Get(db, "user.mockUser1")
	assert.Equal(t, test.MockData["user.mockUser1"], bmap)
	assert.NoError(t, err)

	// success - entry does not exist
	bmap, err = Get(db, "nope")
	assert.Nil(t, bmap)
	assert.NoError(t, err)
}

func TestJoin(t *testing.T) {
	// success - two elements
	name := Join("head", "kind")
	assert.Equal(t, "head.kind", name)

	// success - three elements
	name = Join("head", "kind", "elem")
	assert.Equal(t, "head.kind.elem", name)
}

func TestList(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	names, err := List(db, "pair.mockUser1")
	assert.Equal(t, []string{
		"pair.mockUser1.alpha",
		"pair.mockUser1.bravo",
	}, names)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	err := Set(db, "name", map[string]string{"attr": "data"})
	assert.NoError(t, err)

	// confirm - database
	test.Try(t, db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket([]byte("name"))
		data := string(buck.Get([]byte("attr")))
		assert.Equal(t, "data", data)
		return nil
	}))
}
