package numino

import "fmt"

// counter keeps track of time quantums.
type counter struct {
	// Ticks is the number of ticks that make up a time quantum.
	Ticks int
	// The last recorded time quantum
	lastQuantum int
}

// Update updates this counter.
//
// Returns true iff the specified number of Ticks have passed between the last
// time quantum and ticks.
func (c *counter) Update(ticks int) bool {
	var elapsed bool
	if ticks-c.lastQuantum > c.Ticks {
		elapsed = true
		c.lastQuantum = ticks
	}
	fmt.Println("last: ", c.lastQuantum, "cur:", ticks)
	return elapsed
}
