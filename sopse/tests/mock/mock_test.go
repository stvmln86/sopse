package mock

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	// success
	requ := Request("GET", "/", "body")
	assert.Equal(t, "GET", requ.Method)
	assert.Equal(t, "/", requ.URL.Path)

	// confirm - body
	bytes, err := io.ReadAll(requ.Body)
	assert.Equal(t, "body", string(bytes))
	assert.NoError(t, err)
}
