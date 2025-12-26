// Package mock implements unit testing mock data and functions.
package mock

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

// Hash is a mock user hash.
var Hash = hash("1.1.1.1", "2026-01-01T12:00:00Z00:00", "salt")

// Files is a map of mock JSON files.
var Files = map[string]map[string]any{
	Hash + ".json": {
		"addr":  "1.1.1.1",
		"init":  "2026-01-01T12:00:00Z00:00",
		"last":  "2026-01-02T18:00:00Z00:00",
		"names": []any{"alpha", "bravo"},
	},

	Hash + ".alpha.json": {
		"body": "Alpha body.\n",
		"hash": hash("Alpha body.\n"),
		"init": "2026-01-01T12:00:00Z00:00",
		"last": "2026-01-01T12:00:00Z00:00",
	},

	Hash + ".bravo.json": {
		"body": "Bravo body.\n",
		"hash": hash("Bravo body.\n"),
		"init": "2026-01-01T12:00:00Z00:00",
		"last": "2026-01-02T18:00:00Z00:00",
	},
}

// hash returns a SHA256 hash from multiple strings.
func hash(elems ...string) string {
	hash := sha256.New()
	for _, elem := range elems {
		hash.Write([]byte(elem))
	}

	return hex.EncodeToString(hash.Sum(nil))
}

// Dire creates and returns a temporary directory containing all mock files.
func Dire(t *testing.T) string {
	dire := t.TempDir()
	for base, pairs := range Files {
		dest := filepath.Join(dire, base)
		bytes, err := json.Marshal(pairs)
		if err != nil {
			t.Fatal(err)
		}

		if err := os.WriteFile(dest, bytes, 0644); err != nil {
			t.Fatal(err)
		}
	}

	return dire
}
