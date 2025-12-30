// Package neat implements data sanitisation and conversion functions.
package neat

import (
	"net"
	"net/http"
	"time"
)

// Addr returns the remote IP address from a Request.
func Addr(r *http.Request) string {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

// Time returns a local Time object from a Unix UTC integer.
func Time(unix int64) time.Time {
	return time.Unix(unix, 0).Local()
}
