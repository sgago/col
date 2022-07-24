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
	"github.com/sgago/col"
	"github.com/sgago/col/slice"
)

// The monotonicity order, that is, always increasing or decreasing.
type Order bool

const (
	// Indicates a monotonic stack where values are always increasing (little to big).
	Increasing Order = true

	// Indicates a monotonic stack where values are always decreasing (big to little).
	Decreasing Order = false
)

// An unbounded, slice-backed monotonic stack data structure with type T elements.
type monostack[T any] struct {
	order Order
	elems []col.PV[T]
}

// Allocates and initializes a new monotonic stack with type T elements.
//
// Elements that would break the monotonic condition are also returned to the caller.
//
// Values are pushed by index in ascending (non-decreasing) order.
// In other wrods, the first value, values[0], will be pushed first.
// The last value, values[len(values)-1] will be pushed last and appear
// on top of the stack.
func New[T any](order Order, cap int, vals ...col.PV[T]) (*monostack[T], []col.PV[T]) {
	s := monostack[T]{
		order: order,
		elems: make([]col.PV[T], 0, cap),
	}

	popped := s.PushMany(vals...)

	return &s, popped
}

// Push adds a value to the top of the stack.
// Elements that would break the monotonic condition are returned to the caller.
func (s *monostack[T]) Push(pv col.PV[T]) []col.PV[T] {
	if s.order {
		return s.pushAsc(pv)
	}

	return s.pushDesc(pv)
}

func (s *monostack[T]) pushAsc(pv col.PV[T]) []col.PV[T] {

	popped := make([]col.PV[T], 0)

	for !s.IsEmpty() && s.Peek().Priority > pv.Priority {
		popped = append(popped, s.Pop())
	}

	s.elems = append(s.elems, pv)

	return popped
}

func (s *monostack[T]) pushDesc(pv col.PV[T]) []col.PV[T] {

	popped := make([]col.PV[T], 0)

	for !s.IsEmpty() && s.Peek().Priority < pv.Priority {
		popped = append(popped, s.Pop())
	}

	s.elems = append(s.elems, pv)

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
func (s *monostack[T]) PushMany(pvs ...col.PV[T]) []col.PV[T] {

	popped := make([]col.PV[T], 0)

	if len(pvs) != 0 {
		for _, value := range pvs {
			popped = append(popped, s.Push(value)...)
		}
	}

	return popped
}

// Pop removes and returns the top element of the monotonic stack.
//
// This method panics if the stack is empty.
func (s *monostack[T]) Pop() col.PV[T] {
	if s.elems == nil || len(s.elems) == 0 {
		panic("The stack is empty.")
	}

	_, last, _ := slice.Last(s.elems, nil)

	s.elems = slice.RemoveLast(s.elems)

	return last
}

// Pop returns the top element of the monotonic stack.
//
// This method panics if the monotonic stack is empty.
func (s *monostack[T]) Peek() col.PV[T] {
	if s.elems == nil || len(s.elems) == 0 {
		panic("The stack is empty.")
	}

	return s.elems[len(s.elems)-1]
}

// Count returns the number of elements in the monotonic stack.
func (s *monostack[T]) Count() int {
	return len(s.elems)
}

// Capacity returns the capacity of the monotonic stack.
func (s *monostack[T]) Capacity() int {
	return cap(s.elems)
}

// IsEmpty returns true if the monotonic stack has no elements;
// otherwise, false.
func (s *monostack[T]) IsEmpty() bool {
	return s.Count() == 0
}

// Clear removes all elements from the monotonic stack.
// It maintains the stack's existing capacity.
func (s *monostack[T]) Clear() {
	s.elems = slice.Clear(s.elems)
}
