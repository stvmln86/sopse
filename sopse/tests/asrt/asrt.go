// Package asrt implements unit testing assertion functions.
package asrt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// unmarshal returns a JSON map from a byte slice.
func unmarshal(t *testing.T, bytes []byte) map[string]any {
	var jmap map[string]any
	if err := json.Unmarshal(bytes, &jmap); err != nil {
		t.Fatal(err)
	}

	return jmap
}

// Error asserts an error's message is equal to a formatted string.
func Error(t *testing.T, err error, text string, elems ...any) bool {
	text = fmt.Sprintf(text, elems...)
	return assert.EqualError(t, err, text)
}

// File asserts a file's body is equal to a JSON map.
func File(t *testing.T, orig string, want map[string]any) bool {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		t.Fatal(err)
	}

	jmap := unmarshal(t, bytes)
	return assert.Equal(t, want, jmap)
}

// FileSubset asserts a file's body contains a JSON map.
func FileSubset(t *testing.T, orig string, want map[string]any) bool {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		t.Fatal(err)
	}

	jmap := unmarshal(t, bytes)
	return assert.Subset(t, jmap, want)
}

// Response asserts a response's body is equal to a JSON map.
func Response(t *testing.T, w *httptest.ResponseRecorder, want map[string]any) bool {
	rslt := w.Result()
	bytes, err := io.ReadAll(rslt.Body)
	if err != nil {
		t.Fatal(err)
	}

	jmap := unmarshal(t, bytes)
	return assert.Equal(t, want, jmap)
}
