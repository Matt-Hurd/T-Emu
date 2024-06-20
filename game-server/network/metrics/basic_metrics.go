package metrics

import (
	"container/ring"
	"sync"
)

type BasicMetrics struct {
	value        float32
	AverageValue float32
	recentValues *ring.Ring
	mu           sync.Mutex
}

func NewBasicMetrics() *BasicMetrics {
	return &BasicMetrics{
		recentValues: ring.New(8),
	}
}

func (g *BasicMetrics) Set(value float32) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.value = value
	g.recentValues.Value = g.value
	g.recentValues = g.recentValues.Next()
	sum := float32(0.0)
	g.recentValues.Do(func(x interface{}) {
		sum += x.(float32)
	})
	g.AverageValue = sum / float32(g.recentValues.Len())
}
