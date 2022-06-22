// Package stack.monotonic provides an unbounded, slice-backed monotonic stack data structure
// or "Monostack".
// "Monotonic" means a collection of elements that is always increasing or decreasing.
// Similar to the "normal" stack, a monotonic stack provides two main operations push and pop.
// Unlike the "normal" stack, however, a monotonic stack's push operation
// will remove any values that would break the monotonic condition.
// That is, any elements that would not always be increasing or decreasing
// are popped off the stack and returned to the caller.
//
//	1. Push - adds an element to the top of the stack. It also removes and returns all elements that would break the monotonic condition.
//	2. Pop - removes and returns the top element of the stack.
//
// This monotonic stack data structure is NOT thread-safe.
package monotonic

import (
	"github.com/sgago/collections"
	"github.com/sgago/collections/slice"
)

// An unbounded, slice-backed monotonic stack data structure with type T elements.
type monostack[T any] struct {
	ascending bool
	elements  []collections.KeyValue[T]
}

// Allocates and initializes a new monotonic stack with type T elements.
//
// Elements that would break the monotonic condition are also returned to the caller.
//
// Values are pushed by index in ascending (non-decreasing) order.
// In other wrods, the first value, values[0], will be pushed first.
// The last value, values[len(values)-1] will be pushed last and appear
// on top of the stack.
func New[T any](capacity int, ascending bool, kvs ...collections.KeyValue[T]) (monostack[T], []collections.KeyValue[T]) {
	s := monostack[T]{
		ascending: ascending,
		elements:  make([]collections.KeyValue[T], 0, capacity),
	}

	popped := s.PushMany(kvs...)

	return s, popped
}

// Push adds a value to the top of the stack.
// Elements that would break the monotonic condition are returned to the caller.
func (s *monostack[T]) Push(kv collections.KeyValue[T]) []collections.KeyValue[T] {
	if s.ascending {
		return s.pushAsc(kv)
	}

	return s.pushDesc(kv)
}

func (s *monostack[T]) pushAsc(kv collections.KeyValue[T]) []collections.KeyValue[T] {

	popped := make([]collections.KeyValue[T], 0)

	for !s.IsEmpty() && s.Peek().Key > kv.Key {
		popped = append(popped, s.Pop())
	}

	s.elements = append(s.elements, kv)

	return popped
}

func (s *monostack[T]) pushDesc(kv collections.KeyValue[T]) []collections.KeyValue[T] {

	popped := make([]collections.KeyValue[T], 0)

	for !s.IsEmpty() && s.Peek().Key < kv.Key {
		popped = append(popped, s.Pop())
	}

	s.elements = append(s.elements, kv)

	return popped
}

// PushMany adds multiple values on top of the monotonic stack.
// Elements that would break the monotonic condition are returned to the caller.
//
// Values are pushed into the stack by index in ascending (non-decreasing) order.
// In other words, the first value, values[0], will be pushed first.
// The last value, values[len(values)-1] will be pushed last and appear on top of the stack.
//
// If no values are supplied, then nothing will be pushed.
func (s *monostack[T]) PushMany(values ...collections.KeyValue[T]) []collections.KeyValue[T] {

	popped := make([]collections.KeyValue[T], 0)

	if len(values) != 0 {
		for _, value := range values {
			popped = append(popped, s.Push(value)...)
		}
	}

	return popped
}

// Pop removes and returns the top element of the monotonic stack.
//
// This method panics if the stack is empty.
func (s *monostack[T]) Pop() collections.KeyValue[T] {
	if s.elements == nil || len(s.elements) == 0 {
		panic("The stack is empty.")
	}

	last := slice.Last(s.elements)

	s.elements = slice.RemoveLast(s.elements)

	return last
}

// Pop returns the top element of the monotonic stack.
//
// This method panics if the monotonic stack is empty.
func (s *monostack[T]) Peek() collections.KeyValue[T] {
	if s.elements == nil || len(s.elements) == 0 {
		panic("The stack is empty.")
	}

	return s.elements[len(s.elements)-1]
}

// Count returns the number of elements in the monotonic stack.
func (s *monostack[T]) Count() int {
	return len(s.elements)
}

// Capacity returns the capacity of the monotonic stack.
func (s *monostack[T]) Capacity() int {
	return cap(s.elements)
}

// IsEmpty returns true if the monotonic stack has no elements;
// otherwise, false.
func (s *monostack[T]) IsEmpty() bool {
	return s.Count() == 0
}

// Clear removes all elements from the monotonic stack.
// It maintains the stack's existing capacity.
func (s *monostack[T]) Clear() {
	s.elements = slice.Clear(s.elements)
}