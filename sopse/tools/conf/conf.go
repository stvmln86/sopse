// Package conf implements command-line configuration globals.
package conf

import (
	"flag"
)

// FlagSet is the default global command-line parser.
var FlagSet = flag.NewFlagSet("sopse", flag.ExitOnError)

// Defined command-line flags.
var (
	// system flags
	FlagAddr = FlagSet.String("addr", "127.0.0.1:8000", "host address")
	FlagPath = FlagSet.String("path", "./sopse.db", "database path")

	// rate limiting flags
	FlagRateBody = FlagSet.Int("rateBody", 65536, "max body size")
	FlagRateHits = FlagSet.Int("rateHits", 100, "max hits per hour")
	FlagRateName = FlagSet.Int("rateName", 64, "max name size")
	FlagRateUser = FlagSet.Int("rateUser", 256, "max keys per user")
)

// Parse parses the package's flags from an argument slice.
func Parse(elems []string) {
	FlagSet.Parse(elems)
}
