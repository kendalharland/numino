package numino

// counter keeps track of time quantums.
type counter struct {
	// Ticks is the number of ticks that make up a time quantum.
	Ticks float64
	// The last recorded time quantum
	lastQuantum float64
}

// Update updates this counter.
//
// Returns true iff the specified number of Ticks have passed between the last
// time quantum and ticks.
func (c *counter) Update(ticks float64) bool {
	var elapsed bool
	if ticks-c.lastQuantum > c.Ticks {
		elapsed = true
		c.lastQuantum = ticks
	}
	return elapsed
}
