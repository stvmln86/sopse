// Package neat implements data sanitisation functions.
package neat

import (
	"net"
	"net/http"
	"strings"

	"github.com/stvmln86/sopse/sopse/tools/flag"
)

// cap returns a string capped to a maximum size.
func cap(text string, size int) string {
	if size > 0 && len(text) > size {
		text = text[:size]
	}

	return text
}

// Addr returns the remote IP address from a Request.
func Addr(r *http.Request) string {
	addr, _, _ := net.SplitHostPort(r.RemoteAddr)
	return addr
}

// Body returns a whitespace-trimmed body string capped to flag.BodyMax.
func Body(body string) string {
	body = cap(body, *flag.BodyMax)
	return strings.TrimSpace(body) + "\n"
}

// Name returns a lowercase name string capped to flag.NameMax.
func Name(name string) string {
	name = cap(name, *flag.NameMax)
	return strings.ToLower(name)
}

// UUID returns a lowercase 16-character UUID string.
func UUID(uuid string) string {
	uuid = cap(uuid, 16)
	return strings.ToLower(uuid)
}

// Value returns a Request path value processed with a string function.
func Value(r *http.Request, name string, sfun func(string) string) string {
	data := r.PathValue(name)
	if sfun != nil {
		return sfun(data)
	}

	return data
}
