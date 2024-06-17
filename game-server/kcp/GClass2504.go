package kcp

import (
	"sync"
)

type GClass2504 struct {
	value        float32
	averageValue float32
	gclass1105_0 *GClass1105[float32]
	mu           sync.Mutex
}

func NewGClass2504() *GClass2504 {
	g1105, _ := NewGClass1105[float32](8)
	return &GClass2504{
		gclass1105_0: g1105,
	}
}

// Set sets the value and updates the average
func (g *GClass2504) Set(value float32) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.value = value
	g.gclass1105_0.PushBack(value)
	g.averageValue = g.gclass1105_0.Sum(func(a, b float32) float32 { return a + b }) / float32(g.gclass1105_0.Count())
}
