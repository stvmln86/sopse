package user

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
	"github.com/stvmln86/sopse/sopse/tests/mock"
)

var mockTime = time.Unix(1000, 0).Local()

func mockUser(t *testing.T) *User {
	db := mock.DB(t)
	user, _ := Get(db, "mockUser1")
	return user
}

func TestNew(t *testing.T) {
	// setup
	db := mock.DB(t)

	// success
	user := New(db, "user.mockUser1", "1.1.1.1", mockTime)
	assert.Equal(t, db, user.DB)
	assert.Equal(t, "user.mockUser1", user.Path)
	assert.Equal(t, "1.1.1.1", user.Addr)
	assert.Equal(t, mockTime, user.Init)
}

func TestCreate(t *testing.T) {
	// setup
	db := mock.DB(t)
	r := httptest.NewRequest("GET", "/", nil)

	// success
	user, err := Create(db, r)
	assert.Equal(t, db, user.DB)
	assert.Regexp(t, `user\.[\w-_]{22}`, user.Path)
	assert.Equal(t, "192.0.2.1", user.Addr)
	assert.WithinDuration(t, time.Now(), user.Init, 5*time.Second)
	assert.NoError(t, err)

	// confirm - database
	asrt.Bucket(t, db, user.Path, user.Map())
}

func TestGet(t *testing.T) {
	// setup
	db := mock.DB(t)

	// success
	user, err := Get(db, "mockUser1")
	assert.Equal(t, db, user.DB)
	assert.Equal(t, "user.mockUser1", user.Path)
	assert.Equal(t, "1.1.1.1", user.Addr)
	assert.Equal(t, mockTime, user.Init)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	user := mockUser(t)

	// success
	err := user.Delete()
	assert.NoError(t, err)

	// confirm - database
	asrt.NoBucket(t, user.DB, "user.mockUser1")
}

func TestGetPair(t *testing.T) {
	// setup
	user := mockUser(t)

	// success
	pair, err := user.GetPair("alpha")
	assert.Equal(t, user.DB, pair.DB)
	assert.Equal(t, "pair.mockUser1.alpha", pair.Path)
	assert.Equal(t, "Alpha.", pair.Body)
	assert.Equal(t, mockTime, pair.Init)
	assert.NoError(t, err)
}

func TestListPairs(t *testing.T) {
	// setup
	user := mockUser(t)

	// success
	paths, err := user.ListPairs()
	assert.Equal(t, []string{
		"pair.mockUser1.alpha",
		"pair.mockUser1.bravo",
	}, paths)
	assert.NoError(t, err)
}

func TestMap(t *testing.T) {
	// setup
	user := mockUser(t)

	// success
	bmap := user.Map()
	assert.Equal(t, mock.Data["user.mockUser1"], bmap)
}

func TestSetPair(t *testing.T) {
	// setup
	user := mockUser(t)

	// success
	pair, err := user.SetPair("name", "body")
	assert.Equal(t, user.DB, pair.DB)
	assert.Equal(t, "pair.mockUser1.name", pair.Path)
	assert.Equal(t, "body", pair.Body)
	assert.WithinDuration(t, time.Now(), pair.Init, 5*time.Second)
	assert.NoError(t, err)

	// confirm - database
	asrt.Bucket(t, user.DB, "pair.mockUser1.name", pair.Map())
}

func TestUUID(t *testing.T) {
	// setup
	user := mockUser(t)

	// success
	uuid := user.UUID()
	assert.Equal(t, "mockUser1", uuid)
}
