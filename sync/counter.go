package counter

import "sync"

// NewCounter initializes the counter
func NewCounter() *Counter {
	return &Counter{}
}

// Counter struct to maintain the count
type Counter struct {
	mu    sync.Mutex
	count int
}

// Inc increments count by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

// Value returns the current count
func (c *Counter) Value() int {
	return c.count
}
