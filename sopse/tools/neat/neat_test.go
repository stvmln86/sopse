package neat

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mockTime = time.Unix(1000, 0).Local()

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

func TestAddr(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/", nil)

	// success
	addr := Addr(r)
	assert.Equal(t, "192.0.2.1", addr)
}

func TestTime(t *testing.T) {
	// setup
	failTime := time.Unix(0, 0).Local()

	// success - valid time
	tobj := Time("1000")
	assert.Equal(t, mockTime, tobj)

	// failure - invalid time
	tobj = Time("")
	assert.Equal(t, failTime, tobj)
}

func TestUnix(t *testing.T) {
	// success
	unix := Unix(mockTime)
	assert.Equal(t, "1000", unix)
}

func TestUUID(t *testing.T) {
	// success
	uuid := UUID()
	assert.Len(t, uuid, 22)
}
