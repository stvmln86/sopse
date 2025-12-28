// Package pair implements the Pair type and methods.
package pair

// Pair is a single stored key-value pair.
type Pair struct {
	Body string `json:"body"`
	Hash string `json:"hash"`
	Init int64  `json:"init"`
}
