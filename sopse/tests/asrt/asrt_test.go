package asrt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	// setup
	db := sqlx.MustConnect("sqlite3", ":memory:")

	// success
	data := Get(t, db, "select ?", 1)
	assert.Equal(t, int64(1), data)
}

func TestResponse(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	fmt.Fprintf(w, "body")

	// success
	Response(t, w, http.StatusOK, "body")
}
