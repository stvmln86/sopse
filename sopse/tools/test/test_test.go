package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB(t *testing.T) {
	// success
	db := DB(t)
	assert.Contains(t, db.Path(), "TestDB.db")
}

func TestTry(t *testing.T) {
	// success
	Try(t, nil)
}
