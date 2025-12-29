package ware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	// success
	hand := Apply(mockHandler)
	assert.NotNil(t, hand)
}
