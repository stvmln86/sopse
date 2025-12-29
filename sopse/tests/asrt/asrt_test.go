package asrt

import (
	"testing"

	"github.com/stvmln86/sopse/sopse/tests/mock"
)

func TestBucket(t *testing.T) {
	// setup
	db := mock.DB(t)

	// success
	Bucket(t, db, "user.mockUser1", mock.Data["user.mockUser1"])
}

func TestNoBucket(t *testing.T) {
	// setup
	db := mock.DB(t)

	// success
	NoBucket(t, db, "nope")
}
