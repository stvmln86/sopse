package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/mock"
)

func TestTrim(t *testing.T) {
	// success - below size
	text := trim("text", 5)
	assert.Equal(t, "text", text)

	// success - at size
	text = trim("text", 4)
	assert.Equal(t, "text", text)

	// success - above size
	text = trim("text", 3)
	assert.Equal(t, "tex", text)

	// success - zero size
	text = trim("text", 0)
	assert.Equal(t, "text", text)
}

func TestAddr(t *testing.T) {
	// setup
	r := mock.Request("GET", "/", "")

	// success
	addr := Addr(r)
	assert.Equal(t, "192.0.2.1", addr)
}

func TestBody(t *testing.T) {
	// success
	body := Body("\tBody.\n")
	assert.Equal(t, "Body.\n", body)
}

func TestName(t *testing.T) {
	// success
	name := Name("NAME")
	assert.Equal(t, "name", name)
}

func TestTime(t *testing.T) {
	// setup
	want := time.Unix(0, 0).Local()

	// success
	tobj := Time(0)
	assert.Equal(t, want, tobj)
}

func TestUUID(t *testing.T) {
	// success
	uuid := UUID("AAAABBBBCCCCDDDD...")
	assert.Equal(t, "aaaabbbbccccdddd", uuid)
}
