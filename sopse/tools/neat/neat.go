// Package neat implements data sanitisation and conversion functions.
package neat

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"time"
	"unicode"
)

// Body returns a whitespace-trimmed body with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Hash returns a base64-encoded SHA256 hash of a string.
func Hash(text string) string {
	hash := sha256.Sum256([]byte(text))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

// Name returns a lowercase alphanumeric name with dashes.
func Name(name string) string {
	var runes []rune
	for _, rune := range strings.ToLower(name) {
		switch {
		case unicode.IsLetter(rune) || unicode.IsNumber(rune):
			runes = append(runes, rune)
		case unicode.IsSpace(rune) || rune == '-' || rune == '_':
			runes = append(runes, '-')
		}
	}

	return strings.Trim(string(runes), "-")
}

// Time returns a local Time object from a Unix UTC integer.
func Time(unix int64) time.Time {
	return time.Unix(unix, 0).Local()
}
