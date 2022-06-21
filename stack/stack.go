package stack

import (
	"github.com/sgago/collections/slice"
)

type stack[T any] struct {
	items []T
}

func New[T any](capacity int, values ...T) stack[T] {
	s := stack[T]{items: make([]T, 0, capacity)}

	for _, value := range values {
		s.Push(value)
	}

	return s
}

func (s *stack[T]) Push(values ...T) {
	s.items = append(s.items, values...)
}

func (s *stack[T]) Pop() T {
	if s.items == nil || len(s.items) == 0 {
		panic("The stack is empty.")
	}

	last := slice.Last(s.items)

	s.items = slice.RemoveLast(s.items)

	return last
}

func (s *stack[T]) Peek() T {
	if s.items == nil || len(s.items) == 0 {
		panic("The stack is empty.")
	}

	return s.items[len(s.items)-1]
}

func (s *stack[T]) Count() int {
	return len(s.items)
}

func (s *stack[T]) Capacity() int {
	return cap(s.items)
}

func (s *stack[T]) IsEmpty() bool {
	return s.Count() == 0
}

func (s *stack[T]) Clear() {
	s.items = slice.Clear(s.items)
}
