//package giasecond holds functions for computing time on gigaseconds
package gigasecond

// import path for the time package from the standard library
import "time"

// add 1 gigasecond(1000000000) ia a given time
func AddGigasecond(t time.Time) time.Time {
	
    t=t.Add(time.Second*1000000000)
	return t
}
