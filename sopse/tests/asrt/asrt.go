// Package asrt implements unit testing assertion and retrieval functions.
package asrt

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// Get returns the first value from a database query.
func Get(t *testing.T, db *sqlx.DB, code string, elems ...any) any {
	var data any
	if err := db.Get(&data, code, elems...); err != nil {
		t.Fatal(t)
	}

	return data
}

// Response asserts the status code and body from a ResponseRecorder.
func Response(t *testing.T, w *httptest.ResponseRecorder, code int, body string) {
	rslt := w.Result()
	bytes, err := io.ReadAll(rslt.Body)
	assert.Equal(t, code, rslt.StatusCode)
	assert.Equal(t, body, string(bytes))
	assert.NoError(t, err)
}
