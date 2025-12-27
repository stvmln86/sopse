package neat

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCap(t *testing.T) {
	// success - with size
	text := cap("text...", 4)
	assert.Equal(t, "text", text)

	// success - zero size
	text = cap("text", 0)
	assert.Equal(t, "text", text)
}

func TestAddr(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/", nil)

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

func TestUUID(t *testing.T) {
	// success
	uuid := UUID("AAAABBBBCCCCDDDD...")
	assert.Equal(t, "aaaabbbbccccdddd", uuid)
}

func TestValue(t *testing.T) {
	// setup
	r := httptest.NewRequest("GET", "/", nil)
	r.SetPathValue("name", "\tdata\n")

	// success - with function
	data := Value(r, "name", strings.TrimSpace)
	assert.Equal(t, "data", data)

	// success - without function
	data = Value(r, "name", nil)
	assert.Equal(t, "\tdata\n", data)
}
