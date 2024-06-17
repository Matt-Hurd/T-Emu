package kcp

type GClass2501 struct {
	messages *GClass2502
	bytes    *GClass2502
}

// NewGClass2501 initializes a new GClass2501
func NewGClass2501() *GClass2501 {
	return &GClass2501{
		messages: NewGClass2502(),
		bytes:    NewGClass2502(),
	}
}

// Increment increments the message and byte counts
func (g *GClass2501) Increment(value int) {
	g.messages.Increment()
	g.bytes.Increment(value)
}

// Commit commits the messages and bytes
func (g *GClass2501) Commit() {
	g.messages.Commit()
	g.bytes.Commit()
}
