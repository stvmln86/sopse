package neat

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAddr(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/", nil)

	// success
	addr := Addr(r)
	assert.Equal(t, "192.0.2.1", addr)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(1000, 0).Local()

	// success
	tobj := Time(1000)
	assert.Equal(t, want, tobj)
}
