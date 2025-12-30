package page

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
	page, err := Create(db, 1, "Body.\n")
	assert.EqualValues(t, &Page{
		db:   db,
		ID:   int64(6),
		Init: time.Now().Unix(),
		Note: int64(1),
		Body: "Body.\n",
		Hash: "RCYc4kLhuZ1Sx9Kky228taSrUHvtm5swMGKpafr-HR4",
	}, page)
	assert.NoError(t, err)

	// confirm - database
	asrt.Row(t, db, "select body from Pages where id=6", "Body.\n")
}

func TestLatest(t *testing.T) {
	// setup
	db := mock.DB()

	// success
	page, err := Latest(db, 1)
	assert.EqualValues(t, &Page{
		db:   db,
		ID:   int64(2),
		Init: int64(1767272400),
		Note: int64(1),
		Body: "Alpha new.\n",
		Hash: "oE39AcPiuCRhVTo_8oY1KXHsoA12gx1l-Cnvm_-REPY",
	}, page)
	assert.NoError(t, err)

	// failure - page does not exist
	page, err = Latest(db, -1)
	assert.Nil(t, page)
	assert.NoError(t, err)
}

func TestInitTime(t *testing.T) {
	// setup
	page, _ := Latest(mock.DB(), 1)
	want := neat.Time(1767272400)

	// success
	tobj := page.InitTime()
	assert.Equal(t, want, tobj)
}

func TestVerify(t *testing.T) {
	// setup
	page, _ := Latest(mock.DB(), 1)

	// success - true
	okay := page.Verify()
	assert.True(t, okay)

	// setup
	page.Body = ""

	// success - false
	okay = page.Verify()
	assert.False(t, okay)
}
