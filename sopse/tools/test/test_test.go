package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// setup
	db := MockDB(t)

	// success
	data := Get(db, "user.mockUser1", "init")
	assert.Equal(t, "1000", data)
}

func TestMockDB(t *testing.T) {
	// success
	db := MockDB(t)
	assert.NotNil(t, db)

	// confirm - mockData
	for name, pairs := range mockData {
		for attr, want := range pairs {
			data := Get(db, name, attr)
			assert.Equal(t, want, data)
		}
	}
}
