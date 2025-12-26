package asrt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponse(t *testing.T) {
	// setup
	w := httptest.NewRecorder()
	fmt.Fprintf(w, "body")

	// success
	Response(t, w, http.StatusOK, "body")
}
