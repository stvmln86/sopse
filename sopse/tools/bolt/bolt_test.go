package bolt

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
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
	db := test.DB(t, false)
	test.Set(t, db, "name", "attr", "data")

	// success
	err := Delete(db, "name")
	assert.NoError(t, err)

	// confirm - database
	okay := test.Exists(t, db, "name")
	assert.False(t, okay)
}

func TestExists(t *testing.T) {
	// setup
	db := test.DB(t, false)
	test.Set(t, db, "name", "attr", "data")

	// success - true
	okay, err := Exists(db, "name")
	assert.True(t, okay)
	assert.NoError(t, err)

	// success - false
	okay, err = Exists(db, "nope")
	assert.False(t, okay)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.DB(t, false)
	test.Set(t, db, "name", "attr", "data")

	// success
	data, err := Get(db, "name", "attr")
	assert.Equal(t, "data", data)
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
	db := test.DB(t, false)
	test.Set(t, db, "name1", "attr", "data")
	test.Set(t, db, "name2", "attr", "data")

	// success
	names, err := List(db, "name")
	assert.Equal(t, []string{"name1", "name2"}, names)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := test.DB(t, false)

	// success
	err := Set(db, "name", map[string]string{"attr": "data"})
	assert.NoError(t, err)

	// confirm - database
	data := test.Get(t, db, "name", "attr")
	assert.Equal(t, "data", data)
}
