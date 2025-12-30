package test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertError(t *testing.T) {
	// setup
	err := errors.New("error")

	// success
	okay := AssertError(t, err, "%s", "error")
	assert.True(t, okay)
}
