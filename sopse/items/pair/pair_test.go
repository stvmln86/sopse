package pair

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
	"github.com/stvmln86/sopse/sopse/tests/mock"
)

var mockTime = time.Unix(1000, 0).Local()

func mockPair(t *testing.T) *Pair {
	db := mock.DB(t)
	pair, _ := Get(db, "mockUser1", "alpha")
	return pair
}

func TestNew(t *testing.T) {
	// setup
	db := mock.DB(t)

	// success
	pair := New(db, "pair.mockUser1.alpha", "Alpha.", mockTime)
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, "pair.mockUser1.alpha", pair.Path)
	assert.Equal(t, "Alpha.", pair.Body)
	assert.Equal(t, mockTime, pair.Init)
}

func TestGet(t *testing.T) {
	// setup
	db := mock.DB(t)

	// success
	pair, err := Get(db, "mockUser1", "alpha")
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, "pair.mockUser1.alpha", pair.Path)
	assert.Equal(t, "Alpha.", pair.Body)
	assert.Equal(t, mockTime, pair.Init)
	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	// setup
	db := mock.DB(t)

	// success
	pair, err := Set(db, "mockUser1", "alpha", "body")
	assert.Equal(t, db, pair.DB)
	assert.Equal(t, "pair.mockUser1.alpha", pair.Path)
	assert.Equal(t, "body", pair.Body)
	assert.WithinDuration(t, time.Now(), pair.Init, 5*time.Second)
	assert.NoError(t, err)

	// confirm - database
	asrt.Bucket(t, db, "pair.mockUser1.alpha", pair.Map())
}

func TestDelete(t *testing.T) {
	// setup
	pair := mockPair(t)

	// success
	err := pair.Delete()
	assert.NoError(t, err)

	// confirm - database
	asrt.NoBucket(t, pair.DB, "pair.mockUser1.alpha")
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
	assert.Equal(t, mock.Data["pair.mockUser1.alpha"], bmap)
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
	asrt.Bucket(t, pair.DB, "pair.mockUser1.alpha", pair.Map())
}
