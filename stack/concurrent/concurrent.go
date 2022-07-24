// Package concurrent provides an unbounded, slice-backed,
// thread-safe stack data structure.
// A stack data structure provides two main operations:
//
//	1. Push which adds an element to the top of the stack.
//	2. Pop which removes and returns the top element of the stack.
package concurrent

import (
	"sync"

	"github.com/sgago/col/slice"
)

// An unbounded, slice-backed, thread-safe stack data structure with type T elements.
type concstack[T any] struct {
	elements []T
	mu       sync.Mutex
}

// Allocates and initializes a new stack with type T elements.
//
// Values are pushed by index in ascending (non-decreasing) order.
// In other wrods, the first value, values[0], will be pushed first.
// The last value, values[len(values)-1] will be pushed last and appear
// on top of the stack.
func New[T any](capacity int, values ...T) *concstack[T] {
	s := concstack[T]{elements: make([]T, 0, capacity)}

	s.PushMany(values...)

	return &s
}

// Push adds a value to the top of the stack.
func (s *concstack[T]) Push(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.elements = append(s.elements, value)
}

// PushMany adds multiple values on top of the stack.
//
// Values are pushed into the stack by index in ascending (non-decreasing) order.
// In other words, the first value, values[0], will be pushed first.
// The last value, values[len(values)-1] will be pushed last and appear on top of the stack.
//
// If no values are supplied, then nothing will be pushed.
func (s *concstack[T]) PushMany(values ...T) {
	if len(values) != 0 {
		s.mu.Lock()
		defer s.mu.Unlock()

		s.elements = append(s.elements, values...)
	}
}

// Pop removes and returns the top element of the stack.
//
// This method panics if the stack is empty.
func (s *concstack[T]) Pop() T {
	if s.elements == nil || len(s.elements) == 0 {
		panic("The stack is empty.")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	_, last, _ := slice.Last(s.elements, nil)

	s.elements = slice.RemoveLast(s.elements)

	return last
}

// Pop returns the top element of the stack.
//
// This method panics if the stack is empty.
func (s *concstack[T]) Peek() T {
	if s.elements == nil || len(s.elements) == 0 {
		panic("The stack is empty.")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.elements[len(s.elements)-1]
}

// Count returns the number of elements in the stack.
func (s *concstack[T]) Count() int {
	return len(s.elements)
}

// Capacity returns the capacity of the stack.
func (s *concstack[T]) Capacity() int {
	return cap(s.elements)
}

// IsEmpty returns true if the stack has no elements;
// otherwise, false.
func (s *concstack[T]) IsEmpty() bool {
	return s.Count() == 0
}

// Clear removes all elements from the stack.
// It maintains the stack's existing capacity.
func (s *concstack[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.elements = slice.Clear(s.elements)
}
