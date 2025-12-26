package conf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// setup
	elems := []string{
		"-addr", ":1234",
		"-path", "test.db",
		"-rateBody", "1111",
		"-rateHits", "2222",
		"-rateName", "3333",
		"-rateUser", "4444",
	}

	// success
	Parse(elems)
	assert.Equal(t, *FlagAddr, ":1234")
	assert.Equal(t, *FlagPath, "test.db")
	assert.Equal(t, *FlagRateBody, 1111)
	assert.Equal(t, *FlagRateHits, 2222)
	assert.Equal(t, *FlagRateName, 3333)
	assert.Equal(t, *FlagRateUser, 4444)
}
