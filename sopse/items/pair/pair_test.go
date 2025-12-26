package pair

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
	"github.com/stvmln86/sopse/sopse/tests/mock"
)

func mockPair() (*Pair, error) {
	db := mock.DB()
	return Get(db, 1, "alpha")
}

func TestCreate(t *testing.T) {
	// setup
	db := mock.DB()

	// success
	pair, err := Create(db, 1, "name", "Body.\n")
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, int64(3), pair.ID)
	assert.Equal(t, time.Now().Unix(), pair.Init)
	assert.Equal(t, int64(1), pair.User)
	assert.Equal(t, "name", pair.Name)
	assert.Equal(t, "Body.\n", pair.Body)
	assert.NoError(t, err)

	// confirm - database
	name := asrt.Get(t, db, "select name from Pairs where id=3")
	assert.Equal(t, "name", name)
}

func TestGet(t *testing.T) {
	// setup
	db := mock.DB()

	// success - pair exists
	pair, err := Get(db, 1, "alpha")
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, int64(1), pair.ID)
	assert.Equal(t, int64(1000), pair.Init)
	assert.Equal(t, int64(1), pair.User)
	assert.Equal(t, "alpha", pair.Name)
	assert.Equal(t, "Alpha body.\n", pair.Body)
	assert.NoError(t, err)

	// success - pair does not exist
	pair, err = Get(db, 1, "nope")
	assert.Nil(t, pair)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	pair, _ := mockPair()

	// success
	err := pair.Delete()
	assert.NoError(t, err)

	// confirm - database
	size := asrt.Get(t, pair.DB, "select count(*) from Pairs where id=1")
	assert.Zero(t, size)
}

func TestExists(t *testing.T) {
	// setup
	pair, _ := mockPair()

	// success
	ok, err := pair.Exists()
	assert.True(t, ok)
	assert.NoError(t, err)
}
