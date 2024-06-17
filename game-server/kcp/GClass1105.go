package kcp

import (
	"errors"
	"fmt"
	"sync"
)

type GClass1105[T any] struct {
	buffer   []T
	capacity int
	start    int
	end      int
	size     int
	mu       sync.Mutex
}

func NewGClass1105[T any](capacity int, items ...T) (*GClass1105[T], error) {
	if capacity < 1 {
		return nil, errors.New("circular buffer cannot have negative or zero capacity")
	}
	if len(items) > capacity {
		return nil, errors.New("too many items to fit circular buffer")
	}

	buffer := make([]T, capacity)
	copy(buffer, items)

	return &GClass1105[T]{
		buffer:   buffer,
		capacity: capacity,
		start:    0,
		end:      len(items) % capacity,
		size:     len(items),
	}, nil
}

func (g *GClass1105[T]) Capacity() int {
	return g.capacity
}

func (g *GClass1105[T]) IsFull() bool {
	return g.size == g.capacity
}

func (g *GClass1105[T]) IsEmpty() bool {
	return g.size == 0
}

func (g *GClass1105[T]) Size() int {
	return g.size
}

func (g *GClass1105[T]) Front() (T, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsEmpty() {
		var zero T
		return zero, errors.New("cannot access an empty buffer")
	}
	return g.buffer[g.start], nil
}

func (g *GClass1105[T]) Back() (T, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsEmpty() {
		var zero T
		return zero, errors.New("cannot access an empty buffer")
	}
	return g.buffer[(g.end-1+g.capacity)%g.capacity], nil
}

func (g *GClass1105[T]) Get(index int) (T, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("cannot access index %d. Buffer is empty", index)
	}
	if index >= g.size {
		var zero T
		return zero, fmt.Errorf("cannot access index %d. Buffer size is %d", index, g.size)
	}
	return g.buffer[(g.start+index)%g.capacity], nil
}

func (g *GClass1105[T]) Set(index int, value T) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsEmpty() {
		return fmt.Errorf("cannot access index %d. Buffer is empty", index)
	}
	if index >= g.size {
		return fmt.Errorf("cannot access index %d. Buffer size is %d", index, g.size)
	}
	g.buffer[(g.start+index)%g.capacity] = value
	return nil
}

func (g *GClass1105[T]) PushBack(item T) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsFull() {
		g.buffer[g.end] = item
		g.end = (g.end + 1) % g.capacity
		g.start = g.end
	} else {
		g.buffer[g.end] = item
		g.end = (g.end + 1) % g.capacity
		g.size++
	}
}

func (g *GClass1105[T]) PushFront(item T) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsFull() {
		g.start = (g.start - 1 + g.capacity) % g.capacity
		g.end = g.start
		g.buffer[g.start] = item
	} else {
		g.start = (g.start - 1 + g.capacity) % g.capacity
		g.buffer[g.start] = item
		g.size++
	}
}

func (g *GClass1105[T]) PopBack() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsEmpty() {
		return errors.New("cannot take elements from an empty buffer")
	}
	g.end = (g.end - 1 + g.capacity) % g.capacity
	var zero T
	g.buffer[g.end] = zero
	g.size--
	return nil
}

func (g *GClass1105[T]) PopFront() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsEmpty() {
		return errors.New("cannot take elements from an empty buffer")
	}
	var zero T
	g.buffer[g.start] = zero
	g.start = (g.start + 1) % g.capacity
	g.size--
	return nil
}

func (g *GClass1105[T]) Clear() {
	g.mu.Lock()
	defer g.mu.Unlock()

	g.start = 0
	g.end = 0
	g.size = 0
	for i := range g.buffer {
		var zero T
		g.buffer[i] = zero
	}
}

func (g *GClass1105[T]) ToArray() []T {
	g.mu.Lock()
	defer g.mu.Unlock()

	array := make([]T, g.size)
	for i := 0; i < g.size; i++ {
		array[i] = g.buffer[(g.start+i)%g.capacity]
	}
	return array
}

func (g *GClass1105[T]) ToArraySegments() [][]T {
	g.mu.Lock()
	defer g.mu.Unlock()

	if g.IsEmpty() {
		return [][]T{{}}
	}
	if g.start < g.end {
		return [][]T{g.buffer[g.start:g.end]}
	}
	return [][]T{g.buffer[g.start:], g.buffer[:g.end]}
}

// Iterate returns a channel for iterating over the buffer's elements
func (g *GClass1105[T]) Iterate() <-chan T {
	ch := make(chan T)
	go func() {
		g.mu.Lock()
		defer g.mu.Unlock()

		for i := 0; i < g.size; i++ {
			ch <- g.buffer[(g.start+i)%g.capacity]
		}
		close(ch)
	}()
	return ch
}

// Sum returns the sum of the values in the buffer
func (g *GClass1105[T]) Sum(sumFunc func(T, T) T) T {
	g.mu.Lock()
	defer g.mu.Unlock()

	var sum T
	for i := 0; i < g.size; i++ {
		sum = sumFunc(sum, g.buffer[(g.start+i)%g.capacity])
	}
	return sum
}

// Count returns the number of values in the buffer
func (g *GClass1105[T]) Count() int {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.size
}
