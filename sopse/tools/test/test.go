// Package test implements unit testing data and functions.
package test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertFile asserts a file contains a marshalled JSON value.
func AssertFile(t *testing.T, orig string, data any) {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		t.Fatal(err)
	}

	wants, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(t, string(wants), string(bytes))
}

// TempFile returns a temporary file containing a JSON value.
func TempFile(t *testing.T, data any) string {
	dest := filepath.Join(t.TempDir(), t.Name()+".json")
	bytes, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(dest, bytes, 0644); err != nil {
		t.Fatal(err)
	}

	return dest
}
