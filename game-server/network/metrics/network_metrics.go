package metrics

type NetworkMetrics struct {
	Lose                      *LossMetrics
	Rtt                       *BasicMetrics
	Disordered                *AggregateMetrics
	Duplicated                *AggregateMetrics
	Received                  *MessageMetrics
	Sent                      *MessageMetrics
	ReliableReceived          *MessageMetrics
	ReliableSent              *MessageMetrics
	ReliableSegmentalReceived *MessageMetrics
	ReliableSegmentalSent     *MessageMetrics
	UnreliableReceived        *MessageMetrics
	UnreliableSent            *MessageMetrics
	ReceivedQueue             *BasicMetrics
	SentQueue                 *BasicMetrics
}

func NewNetworkMetrics() *NetworkMetrics {
	return &NetworkMetrics{
		Lose:                      NewLossMetrics(),
		Rtt:                       NewBasicMetrics(),
		Disordered:                NewAggregateMetrics(),
		Duplicated:                NewAggregateMetrics(),
		Received:                  NewMessageMetrics(),
		Sent:                      NewMessageMetrics(),
		ReliableReceived:          NewMessageMetrics(),
		ReliableSent:              NewMessageMetrics(),
		ReliableSegmentalReceived: NewMessageMetrics(),
		ReliableSegmentalSent:     NewMessageMetrics(),
		UnreliableReceived:        NewMessageMetrics(),
		UnreliableSent:            NewMessageMetrics(),
		ReceivedQueue:             NewBasicMetrics(),
		SentQueue:                 NewBasicMetrics(),
	}
}

func (g *NetworkMetrics) Commit() {
	g.Lose.Commit()
	g.Received.Commit()
	g.Sent.Commit()
	g.ReliableReceived.Commit()
	g.ReliableSent.Commit()
	g.UnreliableReceived.Commit()
	g.UnreliableSent.Commit()
}
