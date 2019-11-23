// This package calculates the moment when someone has lived for 10^9 seconds
package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	// This function adds 10^9 seconds to whatever time is given

	gigaseconds := time.Duration(1000000000) * time.Second
	return t.Add(gigaseconds)
}
