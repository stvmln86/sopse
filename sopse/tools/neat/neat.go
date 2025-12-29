// Package neat implements data sanitisation and conversion functions.
package neat

import (
	"encoding/base64"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Expired returns true if a Time object is past a duration.
func Expired(tobj time.Time, dura time.Duration) bool {
	return time.Now().After(tobj.Add(dura))
}

// Addr returns the remote IP address from a Request.
func Addr(r *http.Request) string {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

// Time returns a local Time object from a Unix UTC string.
func Time(unix string) time.Time {
	uint, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		return time.Unix(0, 0).Local()
	}

	return time.Unix(uint, 0).Local()
}

// Unix returns a Unix UTC string from a Time object.
func Unix(tobj time.Time) string {
	return strconv.FormatInt(tobj.Unix(), 10)
}

// UUID returns a base64-encoded random UUID.
func UUID() string {
	uuid := uuid.New()
	return base64.RawURLEncoding.EncodeToString(uuid[:])
}
