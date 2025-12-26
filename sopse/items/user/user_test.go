package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
	"github.com/stvmln86/sopse/sopse/tests/mock"
)

func mockUser() (*User, error) {
	db := mock.DB()
	return Get(db, "aaaa")
}

func TestCreate(t *testing.T) {
	// setup
	db := mock.DB()

	// success
	user, err := Create(db, "1.2.3.4")
	assert.Equal(t, db, user.DB)
	assert.Equal(t, int64(3), user.ID)
	assert.Equal(t, "1.2.3.4", user.Addr)
	assert.Equal(t, time.Now().Unix(), user.Init)
	assert.Len(t, user.UUID, 16)
	assert.NoError(t, err)

	// confirm - database
	addr := asrt.Get(t, db, "select addr from Users where id=3")
	assert.Equal(t, "1.2.3.4", addr)
}

func TestGet(t *testing.T) {
	// setup
	db := mock.DB()

	// success - user exists
	user, err := Get(db, "aaaa")
	assert.Equal(t, db, user.DB)
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "1.1.1.1", user.Addr)
	assert.Equal(t, int64(1000), user.Init)
	assert.Equal(t, "aaaa", user.UUID)
	assert.NoError(t, err)

	// success - user does not exist
	user, err = Get(db, "nope")
	assert.Nil(t, user)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	user, _ := mockUser()

	// success
	err := user.Delete()
	assert.NoError(t, err)

	// confirm - database
	size := asrt.Get(t, user.DB, "select count(*) from Users where id=1")
	assert.Zero(t, size)
}

func TestExists(t *testing.T) {
	// setup
	user, _ := mockUser()

	// success
	ok, err := user.Exists()
	assert.True(t, ok)
	assert.NoError(t, err)
}
