package asrt

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	// setup
	orig := filepath.Join(t.TempDir(), t.Name()+".json")
	if err := os.WriteFile(orig, []byte(`{"key": "value"}`), 0644); err != nil {
		t.Fatal(err)
	}

	// success
	okay := File(t, orig, map[string]any{"key": "value"})
	assert.True(t, okay)
}
