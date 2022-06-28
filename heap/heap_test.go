package heap

import (
	"testing"

	"github.com/sgago/col"
	"github.com/stretchr/testify/assert"
)

var pv1 col.PV[int] = col.PV[int]{Priority: 1, Val: 1}
var pv2 col.PV[int] = col.PV[int]{Priority: 2, Val: 2}
var pv3 col.PV[int] = col.PV[int]{Priority: 3, Val: 3}
var pv4 col.PV[int] = col.PV[int]{Priority: 4, Val: 4}

func TestPush_WithMinHeap_ElementsCorrect(t *testing.T) {
	h := New[int](Min, 4)

	h.Push(pv3)
	h.Push(pv2)
	h.Push(pv1)
	h.Push(pv4)

	assert.Equal(t, 4, len(h.elems))
	assert.Equal(t, 1, h.elems[0].Priority)
	assert.Equal(t, 3, h.elems[1].Priority)
	assert.Equal(t, 2, h.elems[2].Priority)
	assert.Equal(t, 4, h.elems[3].Priority)
}

func TestPush_WithMaxHeap_ElementsCorrect(t *testing.T) {
	h := New[int](Max, 4)

	h.Push(pv3)
	h.Push(pv2)
	h.Push(pv1)
	h.Push(pv4)

	assert.Equal(t, 4, len(h.elems))
	assert.Equal(t, 4, h.elems[0].Priority)
	assert.Equal(t, 3, h.elems[1].Priority)
	assert.Equal(t, 1, h.elems[2].Priority)
	assert.Equal(t, 2, h.elems[3].Priority)
}

func TestPop_WithMinHeap_ElementsCorrect(t *testing.T) {
	h := New[int](Min, 4)

	h.Push(pv3)
	h.Push(pv2)
	h.Push(pv1)
	h.Push(pv4)

	values := make([]int, 0, 4)

	values = append(values, h.Pop().Priority)
	values = append(values, h.Pop().Priority)
	values = append(values, h.Pop().Priority)
	values = append(values, h.Pop().Priority)

	assert.Equal(t, 4, len(values))
	assert.Equal(t, 1, values[0])
	assert.Equal(t, 2, values[1])
	assert.Equal(t, 3, values[2])
	assert.Equal(t, 4, values[3])
}

func TestPop_WithMaxHeap_ElementsCorrect(t *testing.T) {
	h := New[int](Max, 4)

	h.Push(pv3)
	h.Push(pv2)
	h.Push(pv1)
	h.Push(pv4)

	values := make([]int, 0, 4)

	values = append(values, h.Pop().Priority)
	values = append(values, h.Pop().Priority)
	values = append(values, h.Pop().Priority)
	values = append(values, h.Pop().Priority)

	assert.Equal(t, 4, len(values))
	assert.Equal(t, 4, values[0])
	assert.Equal(t, 3, values[1])
	assert.Equal(t, 2, values[2])
	assert.Equal(t, 1, values[3])
}

func TestPeek_WithMinHeap_IsCorrect(t *testing.T) {
	h := New[int](Min, 4)

	h.Push(pv3)
	h.Push(pv2)
	h.Push(pv1)
	h.Push(pv4)

	actual := h.Peek().Priority

	assert.Equal(t, 1, actual)
}

func TestPeek_WithMaxHeap_IsCorrect(t *testing.T) {
	h := New[int](Max, 4)

	h.Push(pv3)
	h.Push(pv2)
	h.Push(pv1)
	h.Push(pv4)

	actual := h.Peek().Priority

	assert.Equal(t, 4, actual)
}

func TestIsEmpty_WithElements_IsFalse(t *testing.T) {
	cap := 1

	h := New(Min, cap, col.PV[int]{
		Priority: 1,
		Val:      1,
	})

	assert.False(t, h.IsEmpty())
}

func TestIsEmpty_WithNoElements_IsTrue(t *testing.T) {
	cap := 1

	h := New[int](Min, cap)

	assert.True(t, h.IsEmpty())
}

func TestClear_WithElements_CountIsCorrect(t *testing.T) {
	cap := 1

	h := New(Min, cap, pv1)

	h.Clear()

	assert.True(t, h.IsEmpty())
}

func TestClear_WithElements_CapacityIsSame(t *testing.T) {
	cap := 1

	h := New(Min, cap, pv1)

	h.Clear()

	assert.Equal(t, cap, h.Capacity())
}
