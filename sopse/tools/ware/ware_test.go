package ware

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mockHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "body")
}

func TestApply(t *testing.T) {
	// success
	hand := Apply(mockHandler, 1)
	assert.NotNil(t, hand)
}
