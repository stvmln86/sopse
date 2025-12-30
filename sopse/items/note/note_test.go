package note

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
	"github.com/stvmln86/sopse/sopse/tests/mock"
	"github.com/stvmln86/sopse/sopse/tools/neat"
)

func TestCreate(t *testing.T) {
	// setup
	db := mock.DB()

	// success
	note, err := Create(db, "name")
	assert.EqualValues(t, &Note{
		db:   db,
		ID:   int64(5),
		Init: time.Now().Unix(),
		Flag: OKAY,
		Name: "name",
	}, note)
	assert.NoError(t, err)

	// confirm - database
	asrt.Row(t, db, "select name from Notes where id=5", "name")
}

func TestRead(t *testing.T) {
	// setup
	db := mock.DB()

	// success
	note, err := Read(db, "alpha")
	assert.EqualValues(t, &Note{
		db:   db,
		ID:   int64(1),
		Init: int64(1767268800),
		Flag: OKAY,
		Name: "alpha",
	}, note)
	assert.NoError(t, err)

	// failure - note does not exist
	note, err = Read(db, "nope")
	assert.Nil(t, note)
	assert.Error(t, err)
}

func TestInitTime(t *testing.T) {
	// setup
	note, _ := Read(mock.DB(), "Alpha")
	want := neat.Time(1767268800)

	// success
	tobj := note.InitTime()
	assert.Equal(t, want, tobj)
}

func TestLatest(t *testing.T) {
	// setup
	note, _ := Read(mock.DB(), "alpha")

	// success
	page, err := note.Latest()
	assert.Equal(t, "Alpha new.\n", page.Body)
	assert.NoError(t, err)
}
