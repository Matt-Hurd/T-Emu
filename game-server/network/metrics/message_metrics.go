package metrics

type MessageMetrics struct {
	messages *AggregateMetrics
	bytes    *AggregateMetrics
}

func NewMessageMetrics() *MessageMetrics {
	return &MessageMetrics{
		messages: NewAggregateMetrics(),
		bytes:    NewAggregateMetrics(),
	}
}

func (g *MessageMetrics) Increment(value int) {
	g.messages.Increment()
	g.bytes.Increment(value)
}

func (g *MessageMetrics) Commit() {
	g.messages.Commit()
	g.bytes.Commit()
}
