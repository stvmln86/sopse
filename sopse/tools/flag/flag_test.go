package flag

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

	// confirm - system flags
	assert.Equal(t, *Addr, ":1234")
	assert.Equal(t, *Path, "test.db")

	// confirm - rate limiting flags
	assert.Equal(t, *RateBody, 1111)
	assert.Equal(t, *RateHits, 2222)
	assert.Equal(t, *RateName, 3333)
	assert.Equal(t, *RateUser, 4444)
}
