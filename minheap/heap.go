package heap

import (
	"github.com/sgago/collections"
	"github.com/sgago/collections/slice"
)

// The heap sort order, that is, a min or max heap.
type HeapSort bool

const (
	// Indicates a minheap data structure.
	Min HeapSort = true

	// Indicates a maxheap data structure.
	Max HeapSort = false
)

type heap[T any] struct {
	elems      []collections.KV[T]
	bubbleUp   func(index int)
	bubbleDown func(index int)
}

func New[T any](sort HeapSort, cap int, vals ...collections.KV[T]) *heap[T] {
	h := heap[T]{
		elems: make([]collections.KV[T], 0, cap),
	}

	if sort == Min {
		h.bubbleUp = h.bubbleUpMinHeap
		h.bubbleDown = h.bubbleDownMinHeap
	} else {
		h.bubbleUp = h.bubbleUpMaxHeap
		h.bubbleDown = h.bubbleDownMaxHeap
	}

	for _, val := range vals {
		h.Push(val.Key, val.Val)
	}

	return &h
}

func (h *heap[T]) Push(key int, val T) {
	h.elems = append(
		h.elems,
		collections.KV[T]{Key: key, Val: val})

	h.bubbleUp(len(h.elems) - 1)
}

func (h *heap[T]) Pop() collections.KV[T] {
	if h.elems == nil || len(h.elems) == 0 {
		panic("The minheap is empty.")
	}

	val := h.elems[0]

	h.elems[0], _ = slice.Last(h.elems, nil)
	h.elems = slice.RemoveLast(h.elems)

	h.bubbleDown(0)

	return val
}

func (h *heap[T]) Peek() collections.KV[T] {
	if h.elems == nil || len(h.elems) == 0 {
		panic("The minheap is empty.")
	}

	return collections.KV[T]{Key: h.elems[0].Key, Val: h.elems[0].Val}
}

// Count returns the number of elements in the heap.
func (h *heap[T]) Count() int {
	return len(h.elems)
}

// Capacity returns the capacity of the heap.
func (h *heap[T]) Capacity() int {
	return cap(h.elems)
}

// IsEmpty returns true if the heap has no elements;
// otherwise, false.
func (h *heap[T]) IsEmpty() bool {
	return h.Count() == 0
}

// Clear removes all elements from the heap.
// It maintains the heap's existing capacity.
func (h *heap[T]) Clear() {
	h.elems = slice.Clear(h.elems)
}

func (h *heap[T]) bubbleUpMinHeap(index int) {
	parent := getParentIndex(index)

	if h.elems[parent].Key > h.elems[index].Key {
		h.elems = slice.Swap(h.elems, parent, index)

		if parent > 0 {
			h.bubbleUpMinHeap(parent)
		}
	}
}

func (h *heap[T]) bubbleUpMaxHeap(index int) {
	parent := getParentIndex(index)

	if h.elems[parent].Key < h.elems[index].Key {
		h.elems = slice.Swap(h.elems, parent, index)

		if parent > 0 {
			h.bubbleUpMaxHeap(parent)
		}
	}
}

func (h *heap[T]) bubbleDownMinHeap(index int) {
	if !h.isLeaf(index) {
		left := getLeftChildIndex(index)
		right := getRightChildIndex(index)

		smallest := left
		isRightEmpty := right >= len(h.elems)

		if !isRightEmpty && h.elems[right].Key < h.elems[left].Key {
			smallest = right
		}

		if h.elems[index].Key > h.elems[smallest].Key {
			h.elems = slice.Swap(h.elems, smallest, index)

			h.bubbleDownMinHeap(smallest)
		}
	}
}

func (maxheap *heap[T]) bubbleDownMaxHeap(index int) {
	if !maxheap.isLeaf(index) {
		left := getLeftChildIndex(index)
		right := getRightChildIndex(index)

		largest := left
		isRightEmpty := right >= len(maxheap.elems)

		if !isRightEmpty && maxheap.elems[left].Key > maxheap.elems[right].Key {
			largest = left
		}

		if maxheap.elems[index].Key < maxheap.elems[largest].Key {
			maxheap.elems = slice.Swap(maxheap.elems, largest, index)

			maxheap.bubbleDownMaxHeap(largest)
		}
	}
}

func (h *heap[T]) isLeaf(index int) bool {
	return len(h.elems) <= getLeftChildIndex(index)
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
