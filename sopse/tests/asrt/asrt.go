// Package asrt implements unit testing assertion functions.
package asrt

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// File asserts a file's body is equal to a JSON map.
func File(t *testing.T, orig string, want map[string]any) bool {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		t.Fatal(err)
	}

	var jmap map[string]any
	if err = json.Unmarshal(bytes, &jmap); err != nil {
		t.Fatal(err)
	}

	return assert.Equal(t, want, jmap)
}
