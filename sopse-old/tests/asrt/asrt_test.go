package asrt

import (
	"testing"
	"time"

	"github.com/stvmln86/sopse/sopse/tests/mock"
)

func TestBucket(t *testing.T) {
	// setup
	db := mock.DB(t)

	// success
	Bucket(t, db, "user.mockUser", mock.Data["user.mockUser"])
}

func TestNoBucket(t *testing.T) {
	// setup
	db := mock.DB(t)

	// success
	NoBucket(t, db, "nope")
}

func TestTimeNow(t *testing.T) {
	// setup
	tobj := time.Now().Add(-1 * time.Second)

	// success - true
	TimeNow(t, tobj)
}
