// Package test implements unit testing data and helpers.
package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertError asserts an error is equal to a formatted string.
func AssertError(t *testing.T, err error, text string, elems ...any) bool {
	text = fmt.Sprintf(text, elems...)
	return assert.EqualError(t, err, text)
}
