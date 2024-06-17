package kcp

import (
	"sync"
)

type GClass2502 struct {
	totalValue   int
	currentValue int
	lastValue    int
	averageValue float32
	values       *GClass1105[int]
	mu           sync.Mutex
}

// NewGClass2502 initializes a new GClass2502
func NewGClass2502() *GClass2502 {
	G1105, _ := NewGClass1105[int](8)
	return &GClass2502{
		values: G1105,
	}
}

// Increment increments the counters
func (g *GClass2502) Increment(values ...int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	incrementValue := 1
	if len(values) > 0 {
		incrementValue = values[0]
	}

	g.totalValue += incrementValue
	g.currentValue += incrementValue
}

func sum(a, b int) int {
	return a + b
}

// Commit commits the current value and calculates the average
func (g *GClass2502) Commit() {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.lastValue = g.currentValue
	g.values.PushBack(g.lastValue)
	g.averageValue = float32(g.values.Sum(sum)) / float32(g.values.Count())
	g.currentValue = 0
}
