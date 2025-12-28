package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertFile(t *testing.T) {
	orig := TempFile(t, 123)
	AssertFile(t, orig, 123)
}

func TestTempFile(t *testing.T) {
	orig := TempFile(t, 123)
	AssertFile(t, orig, 123)
}

func TestTempPath(t *testing.T) {
	path := TempPath(t)
	assert.Contains(t, path, "TestTempPath.json")
}
