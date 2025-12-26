package mock

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
)

func TestHash(t *testing.T) {
	// success
	hash := hash("elem")
	assert.Equal(t, "7ddcdc0ab8180b8893fa2aa21023c2054e546d52409d932dbb4fda71225b63ae", hash)
}

func TestDire(t *testing.T) {
	// success
	dire := Dire(t)
	assert.DirExists(t, dire)

	// confirm - files
	for base, pairs := range Files {
		orig := filepath.Join(dire, base)
		asrt.File(t, orig, pairs)
	}
}
