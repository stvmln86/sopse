package pair

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("without init", func(t *testing.T) {
		pair := New("body")
		assert.Equal(t, "body", pair.Body)
		assert.Equal(t, "Iw2DWNyOiJC0xY3utikS7i8gNXrpKlzIYbmOaP4xrLU", pair.Hash)
		assert.Equal(t, time.Now().Unix(), pair.Init)
	})

	t.Run("with init", func(t *testing.T) {
		pair := New("body", 100)
		assert.Equal(t, "body", pair.Body)
		assert.Equal(t, "Iw2DWNyOiJC0xY3utikS7i8gNXrpKlzIYbmOaP4xrLU", pair.Hash)
		assert.Equal(t, int64(100), pair.Init)
	})
}

func TestCheck(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		pair := &Pair{"body", "Iw2DWNyOiJC0xY3utikS7i8gNXrpKlzIYbmOaP4xrLU", 0}
		okay := pair.Check()
		assert.True(t, okay)
	})

	t.Run("invalid", func(t *testing.T) {
		pair := &Pair{"body", "nope", 0}
		okay := pair.Check()
		assert.False(t, okay)
	})
}

func TestExpired(t *testing.T) {
	init := time.Now().Add(-100 * time.Second).Unix()
	pair := &Pair{"", "", init}

	t.Run("fresh", func(t *testing.T) {
		okay := pair.Expired(200)
		assert.False(t, okay)
	})

	t.Run("stale", func(t *testing.T) {
		okay := pair.Expired(99)
		assert.True(t, okay)
	})
}

func TestTime(t *testing.T) {
	pair := &Pair{"", "", 100}
	want := time.Unix(100, 0).Local()
	tobj := pair.Time()
	assert.Equal(t, want, tobj)
}

func TestUpdate(t *testing.T) {
	pair := &Pair{"", "", 0}
	pair.Update("body")
	assert.Equal(t, "body", pair.Body)
	assert.Equal(t, "Iw2DWNyOiJC0xY3utikS7i8gNXrpKlzIYbmOaP4xrLU", pair.Hash)
}
