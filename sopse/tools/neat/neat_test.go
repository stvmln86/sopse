package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExpired(t *testing.T) {
	// setup
	tobj := time.Now().Add(-1 * time.Minute)

	// success - true
	okay := Expired(tobj, 30*time.Second)
	assert.True(t, okay)

	// success - false
	okay = Expired(tobj, 2*time.Minute)
	assert.False(t, okay)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(100, 0).Local()
	fail := time.Unix(0, 0).Local()

	// success - valid time
	tobj := Time("100")
	assert.Equal(t, want, tobj)

	// failure - invalid time
	tobj = Time("")
	assert.Equal(t, fail, tobj)
}
