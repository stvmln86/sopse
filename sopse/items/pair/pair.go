// Package pair implements the Pair type and methods.
package pair

import (
	"time"

	"github.com/stvmln86/sopse/sopse/tools/neat"
)

// Pair is a single stored key-value pair.
type Pair struct {
	Body string `json:"body"`
	Hash string `json:"hash"`
	Init int64  `json:"init"`
}

// New returns a new Pair.
func New(body string, inits ...int64) *Pair {
	if len(inits) == 0 {
		inits = append(inits, time.Now().Unix())
	}

	return &Pair{body, neat.Hash(body), inits[0]}
}

// Check returns true if the Pair's body matches its hash.
func (p *Pair) Check() bool {
	return neat.Hash(p.Body) == p.Hash
}

// Expired returns true if the Pair's creation time is over a limit.
func (p *Pair) Expired(secs int64) bool {
	return neat.Expired(p.Init, secs)
}

// Time returns the Pair's creation time.
func (p *Pair) Time() time.Time {
	return neat.Time(p.Init)
}

// Update overwrites the Pair's body and hash.
func (p *Pair) Update(body string) {
	p.Body = body
	p.Hash = neat.Hash(body)
}
