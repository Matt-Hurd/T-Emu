package kcp

import (
	"sync"
)

type GClass2503 struct {
	totalCount              int
	loseCount               int
	lastLoseCountValue      int
	averageLoseCountValue   int
	lastLosePercentValue    float32
	averageLosePercentValue float32
	gclass1105_0            *GClass1105[int]
	gclass1105_1            *GClass1105[float32]
	mu                      sync.Mutex
}

func NewGClass2503() *GClass2503 {
	g1105_1, _ := NewGClass1105[int](16)
	g1105_2, _ := NewGClass1105[float32](16)
	return &GClass2503{
		gclass1105_0: g1105_1,
		gclass1105_1: g1105_2,
	}
}

// Increment increments the total and lose counts
func (g *GClass2503) Increment(total, lose int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.totalCount += total
	g.loseCount += lose
}

// Commit commits the current values and calculates averages
func (g *GClass2503) Commit() {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.totalCount == 0 {
		return
	}
	g.lastLoseCountValue = g.loseCount
	g.gclass1105_0.PushBack(g.lastLoseCountValue)
	g.averageLoseCountValue = g.gclass1105_0.Sum(func(a, b int) int { return a + b }) / g.gclass1105_0.Count()
	g.lastLosePercentValue = float32(g.loseCount) / float32(g.totalCount)
	g.gclass1105_1.PushBack(g.lastLosePercentValue)
	g.averageLosePercentValue = g.gclass1105_1.Sum(func(a, b float32) float32 { return a + b }) / float32(g.gclass1105_1.Count())
	g.loseCount = 0
	g.totalCount = 0
}
