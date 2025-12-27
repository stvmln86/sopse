package asrt

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	// setup
	bytes := []byte(`{"key": "value"}`)

	// success
	jmap := unmarshal(t, bytes)
	assert.Equal(t, map[string]any{"key": "value"}, jmap)
}

func TestError(t *testing.T) {
	// setup
	err := errors.New("error")

	// success
	okay := Error(t, err, "%s", "error")
	assert.True(t, okay)
}

func TestFile(t *testing.T) {
	// setup
	data := `{"key": "value"}`
	orig := filepath.Join(t.TempDir(), t.Name()+".json")
	if err := os.WriteFile(orig, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	// success
	okay := File(t, orig, map[string]any{"key": "value"})
	assert.True(t, okay)
}

func TestFileSubset(t *testing.T) {
	// setup
	data := `{"key": "value", "extra": "extra"}`
	orig := filepath.Join(t.TempDir(), t.Name()+".json")
	if err := os.WriteFile(orig, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	// success
	okay := FileSubset(t, orig, map[string]any{"key": "value"})
	assert.True(t, okay)
}

func TestResponse(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	fmt.Fprintf(w, `{"key": "value"}`)

	// success
	okay := Response(t, w, map[string]any{"key": "value"})
	assert.True(t, okay)
}
