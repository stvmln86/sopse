package flag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// setup
	elems := []string{
		"-addr", ":1234",
		"-path", "path.db",
		"-bodyMax", "1234",
		"-nameMax", "1234",
	}

	// success
	Parse(elems)
	assert.Equal(t, ":1234", *Addr)
	assert.Equal(t, "path.db", *Path)
	assert.Equal(t, 1234, *BodyMax)
	assert.Equal(t, 1234, *NameMax)
}
