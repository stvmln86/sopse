package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	// success
	db := DB(t, true)
	assert.Contains(t, db.Path(), "TestDB.db")

	// confirm - database
	for name, pairs := range mockData {
		for attr, want := range pairs {
			data := Get(t, db, name, attr)
			assert.Equal(t, want, data)
		}
	}
}

func TestExists(t *testing.T) {
	// setup
	db := DB(t, false)
	Set(t, db, "name", "attr", "data")

	// success - true
	okay := Exists(t, db, "name")
	assert.True(t, okay)

	// success - false
	okay = Exists(t, db, "nope")
	assert.False(t, okay)
}

func TestGet(t *testing.T) {
	// setup
	db := DB(t, false)
	Set(t, db, "name", "attr", "data")

	// success
	data := Get(t, db, "name", "attr")
	assert.Equal(t, "data", data)
}

func TestSet(t *testing.T) {
	// setup
	db := DB(t, false)

	// success
	Set(t, db, "name", "attr", "data")

	// confirm - database
	data := Get(t, db, "name", "attr")
	assert.Equal(t, "data", data)
}

func TestTry(t *testing.T) {
	// success
	Try(t, nil)
}
