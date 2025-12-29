package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
)

func TestDB(t *testing.T) {
	// success
	db := DB(t)
	assert.Contains(t, db.Path(), "TestDB.db")

	// confirm - database
	for path, bmap := range Data {
		asrt.Bucket(t, db, path, bmap)
	}
}
