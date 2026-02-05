// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

import (
	"time"
)

// AddGigasecond takes a time.Time and returns a time.Time 1 Gigasecond in the future
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
