package queue

import (
	"github.com/sgago/collections/slice"
)

type queue[T any] struct {
	items []T
}

func New[T any](capacity int, values ...T) queue[T] {
	q := queue[T]{items: make([]T, 0, capacity)}

	q.EnqueueRange(values)

	return q
}

func (q *queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

func (q *queue[T]) EnqueueRange(values []T) {
	q.items = append(q.items, values...)
}

func (q *queue[T]) Dequeue() T {
	if q.items == nil || len(q.items) == 0 {
		panic("The queue is empty.")
	}

	first := slice.First(q.items)

	q.items = slice.RemoveFirst(q.items)

	return first
}

func (q *queue[T]) Peek() T {
	if q.items == nil || len(q.items) == 0 {
		panic("The queue is empty.")
	}

	return q.items[0]
}

func (q *queue[T]) Count() int {
	return len(q.items)
}

func (q *queue[T]) IsEmpty() bool {
	return q.Count() == 0
}

func (q *queue[T]) Clear() {
	q.items = make([]T, 0)
}
