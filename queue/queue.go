// Package queue provides an unbounded, slice-backed queue data structure.
// A queue data structure provides two main operations:
//
//	1. Enqueue - adds an element to the back of the queue.
//	2. Dequeue - removes and returns the element in the front of the queue.
//
// This queue data structure is NOT thread-safe.
package queue

import (
	"github.com/sgago/col/slice"
)

// An unbounded, slice-backed queue data structure with type T elements.
type queue[T any] struct {
	elements []T
}

// Allocates and initializes a new queue with type T elements.
//
// Values are enqueued by index in ascending (non-decreasing) order.
// In other wrods, the first value, values[0], will be enqueued first and
// appear in the front of the stack.
// The last value, values[len(values)-1], will be enqueued last and appear
// at the back of the queue.
func New[T any](capacity int, values ...T) *queue[T] {
	q := queue[T]{elements: make([]T, 0, capacity)}

	q.EnqueueMany(values...)

	return &q
}

// Enqueue adds a value to the back of the stack.
func (q *queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// EnqueueMany adds enqueues multiple values into the queue.
//
// Values are enqueued by index in ascending (non-decreasing) order.
// In other words, the first value, values[0], will be enqueued first.
// The last value, values[len(values)-1] will be enqueued last and
// appear at the back of the queue.
//
// If no values are supplied, then nothing will be enqueued.
func (q *queue[T]) EnqueueMany(values ...T) {
	q.elements = append(q.elements, values...)
}

// Dequeue removes and returns the first element in the queue.
//
// This method panics if the queue is empty.
func (q *queue[T]) Dequeue() T {
	if q.elements == nil || len(q.elements) == 0 {
		panic("The queue is empty.")
	}

	first, _ := slice.First(q.elements, nil)

	q.elements = slice.RemoveFirst(q.elements)

	return first
}

// Peek returns the top element of the stack.
//
// This method panics if the stack is empty.
func (q *queue[T]) Peek() T {
	if q.elements == nil || len(q.elements) == 0 {
		panic("The queue is empty.")
	}

	return q.elements[0]
}

// Count returns the number of elements in the queue.
func (q *queue[T]) Count() int {
	return len(q.elements)
}

// Capacity returns the capacity of the stack.
func (q *queue[T]) Capacity() int {
	return cap(q.elements)
}

// IsEmpty returns true if the queue has no elements;
// otherwise, false.
func (q *queue[T]) IsEmpty() bool {
	return q.Count() == 0
}

// Clear removes all elements from queue.
// It maintains the queue's existing capacity.
func (q *queue[T]) Clear() {
	q.elements = make([]T, 0)
}
