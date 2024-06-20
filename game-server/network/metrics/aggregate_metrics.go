package metrics

import (
	"container/ring"
	"sync"
)

type AggregateMetrics struct {
	totalValue   int
	currentValue int
	lastValue    int
	averageValue float32
	recentValues *ring.Ring
	mu           sync.Mutex
}

func NewAggregateMetrics() *AggregateMetrics {
	return &AggregateMetrics{
		recentValues: ring.New(8),
	}
}

func (g *AggregateMetrics) Increment(values ...int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	incrementValue := 1
	if len(values) > 0 {
		incrementValue = values[0]
	}

	g.totalValue += incrementValue
	g.currentValue += incrementValue
}

func (g *AggregateMetrics) Commit() {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.lastValue = g.currentValue
	g.recentValues.Value = g.lastValue
	g.recentValues = g.recentValues.Next()
	sum := 0
	g.recentValues.Do(func(x any) {
		if x != nil {
			sum += x.(int)
		}
	})
	g.averageValue = float32(sum) / float32(g.recentValues.Len())
	g.currentValue = 0
}
