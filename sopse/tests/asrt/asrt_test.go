package asrt

import (
	"testing"

	"github.com/stvmln86/sopse/sopse/tools/test"
)

func TestBucket(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	Bucket(t, db, "user.mockUser1", test.MockData["user.mockUser1"])
}

func TestNoBucket(t *testing.T) {
	// setup
	db := test.DB(t)

	// success
	NoBucket(t, db, "nope")
}
