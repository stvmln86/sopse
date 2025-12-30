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

func TestSet(t *testing.T) {
	// setup
	db := test.DB()

	// success - new pair
	pair, err := Set(db, 1, "name", "body")
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
	okay := test.Get(t, db, "select exists (select 1 from Pairs where id=3)")
	assert.NotZero(t, okay)

	// success - existing pair
	pair, err = Set(db, 1, "alpha", "body")
	assert.EqualValues(t, &Pair{
		DB:   db,
		ID:   int64(1),
		Init: int64(1000),
		User: int64(1),
		Name: "alpha",
		Body: "body",
	}, pair)
	assert.NoError(t, err)

	// confirm - database
	body := test.Get(t, db, "select body from Pairs where id=1")
	assert.Equal(t, "body", body)
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
