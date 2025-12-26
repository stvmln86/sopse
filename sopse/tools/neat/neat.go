// Package neat implements data sanitisation functions.
package neat

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/stvmln86/sopse/sopse/tools/flag"
)

// trim returns a string trimmed to a maximum size.
func trim(text string, size int) string {
	if size > 0 && len(text) > size {
		return text[:size]
	}

	return text
}

// Addr returns the remote address from a Request.
func Addr(r *http.Request) string {
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

// Body returns a whitespace-stripped body string with a trailing newline.
func Body(body string) string {
	body = trim(body, *flag.RateBody)
	return strings.TrimSpace(body) + "\n"
}

// Name returns a lowercase name string.
func Name(name string) string {
	name = trim(name, *flag.RateName)
	return strings.ToLower(name)
}

// Time returns a local Time object from a UTC Unix integer.
func Time(unix int64) time.Time {
	return time.Unix(unix, 0).Local()
}

// UUID returns a lowercase UUID string.
func UUID(uuid string) string {
	uuid = trim(uuid, 16)
	return strings.ToLower(uuid)
}
