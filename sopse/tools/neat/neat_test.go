package neat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExpired(t *testing.T) {
	unix := time.Now().Add(-100 * time.Second).Unix()

	t.Run("fresh", func(t *testing.T) {
		okay := Expired(unix, 200)
		assert.False(t, okay)
	})

	t.Run("stale", func(t *testing.T) {
		okay := Expired(unix, 99)
		assert.True(t, okay)
	})
}

func TestHash(t *testing.T) {
	hash := Hash("a", "b", "c")
	assert.Equal(t, "ungWv48Bz-pBQUDeXa4iI7ADYaOWF3qctBD_YfIAFa0", hash)
}

func TestTime(t *testing.T) {
	want := time.Unix(100, 0).Local()
	tobj := Time(100)
	assert.Equal(t, want, tobj)
}
