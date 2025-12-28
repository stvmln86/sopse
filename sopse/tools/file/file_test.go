package file

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tools/test"
)

func TestCreate(t *testing.T) {
	dest := filepath.Join(t.TempDir(), t.Name()+".json")

	t.Run("success", func(t *testing.T) {
		err := Create(dest, 123)
		test.AssertFile(t, dest, 123)
		assert.NoError(t, err)
	})

	t.Run("error - exists", func(t *testing.T) {
		err := Create(dest, 123)
		assert.Error(t, err)
	})
}

func TestDelete(t *testing.T) {
	orig := test.TempFile(t, 123)
	err := Delete(orig)
	assert.NoFileExists(t, orig)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	orig := test.TempFile(t, 123)

	t.Run("true", func(t *testing.T) {
		okay := Exists(orig)
		assert.True(t, okay)
	})

	t.Run("false", func(t *testing.T) {
		okay := Exists("/nope")
		assert.False(t, okay)
	})
}

func TestRead(t *testing.T) {
	var data int
	orig := test.TempFile(t, 123)
	err := Read(orig, &data)
	assert.Equal(t, 123, data)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	orig := test.TempFile(t, 123)

	t.Run("success", func(t *testing.T) {
		err := Update(orig, 456)
		test.AssertFile(t, orig, 456)
		assert.NoError(t, err)
	})

	t.Run("error - does not exist", func(t *testing.T) {
		err := Update("/nope", 456)
		assert.Error(t, err)
	})
}
