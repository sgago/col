package minheap

import (
	"github.com/sgago/collections"
	"github.com/sgago/collections/slice"
)

type minheap[T any] struct {
	items []collections.KeyValue[T]
}

func New[T any](capacity int, values ...collections.KeyValue[T]) *minheap[T] {
	heap := minheap[T]{items: make([]collections.KeyValue[T], 0, capacity)}

	for _, value := range values {
		heap.Push(value.Key, value.Value)
	}

	return &heap
}

func (heap *minheap[T]) Push(key int, value T) {
	heap.items = append(
		heap.items,
		collections.KeyValue[T]{Key: key, Value: value})

	heap.bubbleUp(len(heap.items) - 1)
}

func (heap *minheap[T]) Pop() collections.KeyValue[T] {
	if heap.items == nil || len(heap.items) == 0 {
		panic("The minheap is empty.")
	}

	min := heap.items[0]

	heap.items[0] = heap.items[len(heap.items)-1]

	heap.bubbleDown(0)

	heap.items = heap.items[:len(heap.items)-1]

	return min
}

func (heap *minheap[T]) Peek() collections.KeyValue[T] {
	if heap.items == nil || len(heap.items) == 0 {
		panic("The minheap is empty.")
	}

	return collections.KeyValue[T]{Key: heap.items[0].Key, Value: heap.items[0].Value}
}

func (heap *minheap[T]) bubbleUp(index int) {
	parent := getParentIndex(index)

	if heap.items[parent].Key > heap.items[index].Key {
		heap.items = slice.Swap(heap.items, parent, index)

		if parent > 0 {
			heap.bubbleUp(parent)
		}
	}
}

func (heap *minheap[T]) bubbleDown(index int) {
	if !heap.isLeaf(index) {
		left := getLeftChildIndex(index)
		right := getRightChildIndex(index)

		smallest := right

		if heap.isLeaf(right) || heap.items[left].Key < heap.items[right].Key {
			smallest = left
		}

		if heap.items[index].Key > heap.items[smallest].Key {
			heap.items = slice.Swap(heap.items, smallest, index)

			heap.bubbleDown(smallest)
		}
	}
}

func getParentIndex(index int) int {
	return (index - 1) / 2
}

func getLeftChildIndex(index int) int {
	return 2*index + 1
}

func getRightChildIndex(index int) int {
	return 2*index + 2
}

func (heap *minheap[T]) isLeaf(index int) bool {
	return len(heap.items) <= getLeftChildIndex(index)
}
