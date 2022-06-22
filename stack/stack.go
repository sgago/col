// Package stack provides an unbounded, slice-backed stack data structure.
// A stack data structure provides two main operations:
//
//	1. Push which adds an element to the top of the stack.
//	2. Pop which removes and returns the top element of the stack.
//
// This stack data structure is NOT thread-safe.
package stack

import (
	"github.com/sgago/collections/slice"
)

// An unbounded, slice-backed stack data structure with type T elements.
type stack[T any] struct {
	elements []T
}

// Allocates and initializes a new stack with type T elements.
//
// Values are pushed by index in ascending (non-decreasing) order.
// In other wrods, the first value, values[0], will be pushed first.
// The last value, values[len(values)-1] will be pushed last and appear
// on top of the stack.
func New[T any](capacity int, values ...T) stack[T] {
	s := stack[T]{elements: make([]T, 0, capacity)}

	s.PushMany(values...)

	return s
}

// Push adds a value to the top of the stack.
func (s *stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// PushMany adds pushes multiple values onto the stack.
//
// Values are pushed into the stack by index in ascending (non-decreasing) order.
// In other words, the first value, values[0], will be pushed first.
// The last value, values[len(values)-1] will be pushed last and appear on top of the stack.
//
// If no values are supplied, then nothing will be pushed.
func (s *stack[T]) PushMany(values ...T) {
	if len(values) != 0 {
		s.elements = append(s.elements, values...)
	}
}

// Pop removes and returns the top element of the stack.
//
// This method panics if the stack is empty.
func (s *stack[T]) Pop() T {
	if s.elements == nil || len(s.elements) == 0 {
		panic("The stack is empty.")
	}

	last := slice.Last(s.elements)

	s.elements = slice.RemoveLast(s.elements)

	return last
}

// Pop returns the top element of the stack.
//
// This method panics if the stack is empty.
func (s *stack[T]) Peek() T {
	if s.elements == nil || len(s.elements) == 0 {
		panic("The stack is empty.")
	}

	return s.elements[len(s.elements)-1]
}

// Count returns the number of elements in the stack.
func (s *stack[T]) Count() int {
	return len(s.elements)
}

// Capacity returns the capacity of the stack.
func (s *stack[T]) Capacity() int {
	return cap(s.elements)
}

// IsEmpty returns true if the stack has no elements;
// otherwise, false.
func (s *stack[T]) IsEmpty() bool {
	return s.Count() == 0
}

// Clear removes all elements from stack.
// It maintains the stack's existing capacity.
func (s *stack[T]) Clear() {
	s.elements = slice.Clear(s.elements)
}
