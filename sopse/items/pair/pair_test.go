package pair

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func mockPair(t *testing.T) *Pair {
	db := test.MockDB(t)
	pair, err := Get(db, 1, "alpha")
	if err != nil {
		t.Fatal(err)
	}

	return pair
}

func TestSet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success - new pair
	pair, err := Set(db, 1, "zulu", "Zulu.\n")
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, int64(3), pair.ID)
	assert.Equal(t, "Zulu.\n", pair.Body)
	assert.Equal(t, time.Now().Unix(), pair.Init)
	assert.Equal(t, "zulu", pair.Name)
	assert.Equal(t, int64(1), pair.User)
	assert.NoError(t, err)

	// confirm - database
	body := test.Get(t, pair.DB, "select body from Pairs where id=3")
	assert.Equal(t, "Zulu.\n", body)

	// success - existing pair
	pair, err = Set(db, 1, "zulu", "Zulu two.\n")
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, int64(3), pair.ID)
	assert.Equal(t, "Zulu two.\n", pair.Body)
	assert.Equal(t, time.Now().Unix(), pair.Init)
	assert.Equal(t, "zulu", pair.Name)
	assert.Equal(t, int64(1), pair.User)
	assert.NoError(t, err)

	// confirm - database
	body = test.Get(t, pair.DB, "select body from Pairs where id=3")
	assert.Equal(t, "Zulu two.\n", body)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success - pair exists
	pair, err := Get(db, 1, "alpha")
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, int64(1), pair.ID)
	assert.Equal(t, "Alpha.\n", pair.Body)
	assert.Equal(t, int64(1000), pair.Init)
	assert.Equal(t, "alpha", pair.Name)
	assert.Equal(t, int64(1), pair.User)
	assert.NoError(t, err)

	// success - pair does not exist
	pair, err = Get(db, -1, "nope")
	assert.Nil(t, pair)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	pair := mockPair(t)

	// success
	err := pair.Delete()
	assert.NoError(t, err)

	// confirm - database
	size := test.Get(t, pair.DB, "select count(*) from Pairs where id=1")
	assert.Zero(t, size)
}

func TestExists(t *testing.T) {
	// setup
	pair := mockPair(t)

	// success
	okay, err := pair.Exists()
	assert.True(t, okay)
	assert.NoError(t, err)
}
