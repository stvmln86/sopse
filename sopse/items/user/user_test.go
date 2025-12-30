package user

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	db := test.DB()
	r := httptest.NewRequest("GET", "/", nil)

	// success
	user, err := Create(db, r)
	assert.EqualValues(t, &User{
		DB:   db,
		ID:   int64(2),
		Init: time.Now().Unix(),
		UUID: user.UUID,
		Addr: "192.0.2.1",
	}, user)
	assert.Len(t, user.UUID, 32)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	// setup
	db := test.DB()

	// success - user exists
	user, err := Get(db, "mockUser")
	assert.EqualValues(t, &User{
		DB:   db,
		ID:   int64(1),
		Init: int64(1000),
		UUID: "mockUser",
		Addr: "1.1.1.1",
	}, user)
	assert.NoError(t, err)

	// success - user does not exist
	user, err = Get(db, "nope")
	assert.Nil(t, user)
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	// setup
	db := test.DB()
	user, _ := Get(db, "mockUser")

	// success
	err := user.Delete()
	assert.NoError(t, err)

	// confirm - database
	okay := test.Get(t, db, "select exists (select 1 from Users where id=1)")
	assert.Zero(t, okay)
}

func TestGetPair(t *testing.T) {
	// setup
	db := test.DB()
	user, _ := Get(db, "mockUser")

	// success
	pair, err := user.GetPair("alpha")
	assert.Equal(t, "alpha", pair.Name)
	assert.NoError(t, err)
}

func TestListPairs(t *testing.T) {
	// setup
	db := test.DB()
	user, _ := Get(db, "mockUser")

	// success
	names, err := user.ListPairs()
	assert.Equal(t, []string{"alpha", "bravo"}, names)
	assert.NoError(t, err)
}

func TestSetPair(t *testing.T) {
	// setup
	db := test.DB()
	user, _ := Get(db, "mockUser")

	// success
	pair, err := user.SetPair("name", "body")
	assert.Equal(t, "name", pair.Name)
	assert.NoError(t, err)

	// confirm - database
	name := test.Get(t, db, "select name from Pairs where id=3")
	assert.Equal(t, "name", name)
}
