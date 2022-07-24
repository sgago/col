package heap

import (
	"github.com/sgago/col"
	"github.com/sgago/col/slice"
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
	elems      []col.PV[T]
	bubbleUp   func(index int)
	bubbleDown func(index int)
}

func New[T any](sort HeapSort, cap int, pvs ...col.PV[T]) *heap[T] {
	h := heap[T]{
		elems: make([]col.PV[T], 0, cap),
	}

	if sort == Min {
		h.bubbleUp = h.bubbleUpMinHeap
		h.bubbleDown = h.bubbleDownMinHeap
	} else {
		h.bubbleUp = h.bubbleUpMaxHeap
		h.bubbleDown = h.bubbleDownMaxHeap
	}

	for _, pv := range pvs {
		h.Push(pv)
	}

	return &h
}

func (h *heap[T]) Push(pv col.PV[T]) {
	h.elems = append(h.elems, pv)

	h.bubbleUp(len(h.elems) - 1)
}

func (h *heap[T]) Pop() col.PV[T] {
	if h.elems == nil || len(h.elems) == 0 {
		panic("The heap is empty.")
	}

	val := h.elems[0]

	_, h.elems[0], _ = slice.Last(h.elems, nil)
	h.elems = slice.RemoveLast(h.elems)

	h.bubbleDown(0)

	return val
}

func (h *heap[T]) Peek() col.PV[T] {
	if h.elems == nil || len(h.elems) == 0 {
		panic("The heap is empty.")
	}

	return col.PV[T]{Priority: h.elems[0].Priority, Val: h.elems[0].Val}
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

	if h.elems[parent].Priority > h.elems[index].Priority {
		h.elems = slice.Swap(h.elems, parent, index)

		if parent > 0 {
			h.bubbleUpMinHeap(parent)
		}
	}
}

func (h *heap[T]) bubbleUpMaxHeap(index int) {
	parent := getParentIndex(index)

	if h.elems[parent].Priority < h.elems[index].Priority {
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

		if !isRightEmpty && h.elems[right].Priority < h.elems[left].Priority {
			smallest = right
		}

		if h.elems[index].Priority > h.elems[smallest].Priority {
			h.elems = slice.Swap(h.elems, smallest, index)

			h.bubbleDownMinHeap(smallest)
		}
	}
}

func (h *heap[T]) bubbleDownMaxHeap(index int) {
	if !h.isLeaf(index) {
		left := getLeftChildIndex(index)
		right := getRightChildIndex(index)

		largest := left
		isRightEmpty := right >= len(h.elems)

		if !isRightEmpty && h.elems[left].Priority > h.elems[right].Priority {
			largest = left
		}

		if h.elems[index].Priority < h.elems[largest].Priority {
			h.elems = slice.Swap(h.elems, largest, index)

			h.bubbleDownMaxHeap(largest)
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
