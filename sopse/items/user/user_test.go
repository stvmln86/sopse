package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func mockUser(t *testing.T) *User {
	db := test.MockDB(t)
	user, err := Get(db, "1111")
	if err != nil {
		t.Fatal(err)
	}

	return user
}

func TestCreate(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success
	user, err := Create(db, "9.9.9.9")
	assert.Equal(t, db, user.DB)
	assert.Equal(t, int64(2), user.ID)
	assert.Equal(t, "9.9.9.9", user.Addr)
	assert.Equal(t, time.Now().Unix(), user.Init)
	assert.Len(t, user.UUID, 16)
	assert.NoError(t, err)

	// confirm - database
	addr := test.Get(t, user.DB, "select addr from Users where id=2")
	assert.Equal(t, "9.9.9.9", addr)
}

func TestGet(t *testing.T) {
	// setup
	db := test.MockDB(t)

	// success - user exists
	user, err := Get(db, "1111")
	assert.Equal(t, db, user.DB)
	assert.Equal(t, int64(1), user.ID)
	assert.Equal(t, "1.1.1.1", user.Addr)
	assert.Equal(t, int64(1000), user.Init)
	assert.Equal(t, "1111", user.UUID)
	assert.NoError(t, err)

	// success - user does not exist
	user, err = Get(db, "nope")
	assert.Nil(t, user)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	user := mockUser(t)

	// success
	err := user.Delete()
	assert.NoError(t, err)

	// confirm - database
	size := test.Get(t, user.DB, "select count(*) from Users where id=1")
	assert.Zero(t, size)
}

func TestExists(t *testing.T) {
	// setup
	user := mockUser(t)

	// success
	okay, err := user.Exists()
	assert.True(t, okay)
	assert.NoError(t, err)
}
