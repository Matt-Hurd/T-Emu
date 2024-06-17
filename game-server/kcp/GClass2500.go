package kcp

type GClass2500 struct {
	lose                      *GClass2503
	rtt                       *GClass2504
	disordered                *GClass2502
	duplicated                *GClass2502
	received                  *GClass2501
	sent                      *GClass2501
	reliableReceived          *GClass2501
	reliableSent              *GClass2501
	reliableSegmentalReceived *GClass2501
	reliableSegmentalSent     *GClass2501
	unreliableReceived        *GClass2501
	unreliableSent            *GClass2501
	receivedQueue             *GClass2504
	sentQueue                 *GClass2504
}

// NewGClass2500 initializes a new GClass2500
func NewGClass2500() *GClass2500 {
	return &GClass2500{
		lose:                      NewGClass2503(),
		rtt:                       NewGClass2504(),
		disordered:                NewGClass2502(),
		duplicated:                NewGClass2502(),
		received:                  NewGClass2501(),
		sent:                      NewGClass2501(),
		reliableReceived:          NewGClass2501(),
		reliableSent:              NewGClass2501(),
		reliableSegmentalReceived: NewGClass2501(),
		reliableSegmentalSent:     NewGClass2501(),
		unreliableReceived:        NewGClass2501(),
		unreliableSent:            NewGClass2501(),
		receivedQueue:             NewGClass2504(),
		sentQueue:                 NewGClass2504(),
	}
}

// Commit commits all metrics
func (g *GClass2500) Commit() {
	g.lose.Commit()
	g.received.Commit()
	g.sent.Commit()
	g.reliableReceived.Commit()
	g.reliableSent.Commit()
	g.unreliableReceived.Commit()
	g.unreliableSent.Commit()
}
