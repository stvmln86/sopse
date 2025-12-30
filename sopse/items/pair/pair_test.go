package pair

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func TestGet(t *testing.T) {
	// setup
	db := test.DB()

	// success - pair exists
	pair, err := Get(db, 1, "alpha")
	assert.EqualValues(t, &Pair{
		DB:   db,
		ID:   int64(1),
		Init: int64(1000),
		User: int64(1),
		Name: "alpha",
		Body: "Alpha.",
	}, pair)
	assert.NoError(t, err)

	// success - pair does not exist
	pair, err = Get(db, -1, "nope")
	assert.Nil(t, pair)
	assert.NoError(t, err)
}

func TestCreate(t *testing.T) {
	// setup
	db := test.DB()

	// success
	pair, err := Create(db, 1, "name", "body")
	assert.EqualValues(t, &Pair{
		DB:   db,
		ID:   int64(3),
		Init: time.Now().Unix(),
		User: int64(1),
		Name: "name",
		Body: "body",
	}, pair)
	assert.NoError(t, err)

	// confirm - database
	name := test.Get(t, db, "select name from Pairs where id=3")
	assert.Equal(t, "name", name)
}

func TestDelete(t *testing.T) {
	// setup
	db := test.DB()
	pair, _ := Get(db, 1, "alpha")

	// success
	err := pair.Delete()
	assert.NoError(t, err)

	// confirm - database
	okay := test.Get(t, db, "select exists (select 1 from Pairs where id=1)")
	assert.Zero(t, okay)
}

func TestUpdate(t *testing.T) {
	// setup
	db := test.DB()
	pair, _ := Get(db, 1, "alpha")

	// success
	err := pair.Update("body")
	assert.NoError(t, err)

	// confirm - database
	body := test.Get(t, db, "select body from Pairs where id=1")
	assert.Equal(t, "body", body)
}
