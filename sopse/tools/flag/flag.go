// Package flag implements command-line configuration globals.
package flag

import (
	"flag"
)

// flagSet is the default global command-line parser.
var flagSet = flag.NewFlagSet("sopse", flag.ExitOnError)

// Defined command-line flags.
var (
	// system flags
	Addr = flagSet.String("addr", "127.0.0.1:8000", "host address")
	Path = flagSet.String("path", "./sopse.db", "database path")

	// rate limiting flags
	RateBody = flagSet.Int("rateBody", 65536, "max body size")
	RateHits = flagSet.Int("rateHits", 100, "max hits per hour")
	RateName = flagSet.Int("rateName", 64, "max name size")
	RateUser = flagSet.Int("rateUser", 256, "max keys per user")
)

// Parse parses the package's flags from an argument slice.
func Parse(elems []string) {
	flagSet.Parse(elems)
}
