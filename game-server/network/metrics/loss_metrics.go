package metrics

import (
	"container/ring"
	"sync"
)

type LossMetrics struct {
	totalCount              int
	loseCount               int
	lastLoseCountValue      int
	AverageLoseCountValue   int
	lastLosePercentValue    float32
	AverageLosePercentValue float32
	recentLoseCounts        *ring.Ring
	recentLosePercents      *ring.Ring
	mu                      sync.Mutex
}

func NewLossMetrics() *LossMetrics {
	return &LossMetrics{
		recentLoseCounts:   ring.New(16),
		recentLosePercents: ring.New(16),
	}
}

func (g *LossMetrics) Increment(total, lose int) {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.totalCount += total
	g.loseCount += lose
}

func (g *LossMetrics) Commit() {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.totalCount == 0 {
		return
	}
	g.lastLoseCountValue = g.loseCount
	g.recentLoseCounts.Value = g.lastLoseCountValue
	g.recentLoseCounts = g.recentLoseCounts.Next()
	sum := 0
	g.recentLoseCounts.Do(func(x interface{}) {
		sum += x.(int)
	})
	g.AverageLoseCountValue = sum / g.recentLoseCounts.Len()
	g.lastLosePercentValue = float32(g.loseCount) / float32(g.totalCount)

	g.recentLosePercents.Value = g.lastLosePercentValue
	g.recentLosePercents = g.recentLosePercents.Next()
	fSum := float32(0.0)
	g.recentLoseCounts.Do(func(x interface{}) {
		fSum += x.(float32)
	})
	g.AverageLosePercentValue = float32(sum) / float32(g.recentLoseCounts.Len())
	g.loseCount = 0
	g.totalCount = 0
}
