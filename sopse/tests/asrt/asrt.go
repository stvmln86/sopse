// Package asrt implements unit testing assertion functions.
package asrt

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

// Error asserts an error is equal to a formatted string.
func Error(t *testing.T, err error, text string, elems ...any) bool {
	text = fmt.Sprintf(text, elems...)
	return assert.EqualError(t, err, text)
}

// Row asserts a database result is equal to a value.
func Row(t *testing.T, db *sqlx.DB, code string, want any) bool {
	var data any
	if err := db.Get(&data, code); err != nil {
		t.Fatal(err)
	}

	switch want := want.(type) {
	case int:
		return assert.Equal(t, int64(want), data)
	default:
		return assert.Equal(t, want, data)
	}
}
