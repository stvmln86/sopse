// Package neat implements data sanitisation and conversion functions.
package neat

import (
	"strconv"
	"time"
)

// Expired returns true if a Time object is past a duration.
func Expired(tobj time.Time, dura time.Duration) bool {
	return time.Now().After(tobj.Add(dura))
}

// Time returns a local Time object from a Unix UTC string.
func Time(unix string) time.Time {
	uint, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		return time.Unix(0, 0).Local()
	}

	return time.Unix(uint, 0).Local()
}
