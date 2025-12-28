// Package neat implements data sanitisation and conversion functions.
package neat

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
)

// Expired returns true if a Unix UTC integer is over a limit.
func Expired(unix, secs int64) bool {
	return time.Now().Unix() > unix+secs
}

// Hash returns a base64-encoded SHA256 hash from joined strings.
func Hash(elems ...string) string {
	hash := sha256.New()
	for _, elem := range elems {
		hash.Write([]byte(elem))
	}

	return base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
}

// Time returns a local Time object from a Unix UTC integer.
func Time(unix int64) time.Time {
	return time.Unix(unix, 0).Local()
}
