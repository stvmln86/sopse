package mock

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sopse/sopse/tests/asrt"
)

func TestDB(t *testing.T) {
	// success
	db := DB()
	assert.NotNil(t, db)

	// confirm - inserts
	size := asrt.Get(t, db, "select count(*) from Users")
	assert.NotZero(t, size)
}

func TestRequest(t *testing.T) {
	// success
	r := Request("GET", "/", "body")
	assert.Equal(t, "GET", r.Method)
	assert.Equal(t, "/", r.URL.Path)

	// confirm - body
	bytes, err := io.ReadAll(r.Body)
	assert.Equal(t, "body", string(bytes))
	assert.NoError(t, err)
}
