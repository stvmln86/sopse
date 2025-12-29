package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	// success
	db := DB(t)
	assert.NotNil(t, db)

	// confirm - mockData
	for name, pairs := range mockData {
		for attr, want := range pairs {
			data := Get(t, db, name, attr)
			assert.Equal(t, want, data)
		}
	}
}

func TestGet(t *testing.T) {
	// setup
	db := DB(t)

	// success
	data := Get(t, db, "user.mockUser1", "init")
	assert.Equal(t, "1000", data)
}

func TestSet(t *testing.T) {
	// setup
	db := DB(t)

	// success
	Set(t, db, "name", "attr", "data")
	data := Get(t, db, "name", "attr")
	assert.Equal(t, "data", data)
}

func TestTry(t *testing.T) {
	// success
	Try(t, nil)
}
