package conf

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	// success
	conf := Parse([]string{"-addr", ":1234"})
	assert.Equal(t, ":1234", conf.Addr)
	assert.Equal(t, "./sopse.db", conf.Dbse)
	assert.Equal(t, int64(4096), conf.BodySize)
	assert.Equal(t, 24*7*time.Hour, conf.PairLife)
	assert.Equal(t, int64(1000), conf.UserRate)
	assert.Equal(t, int64(256), conf.UserSize)
}
