package monotonic

import (
	"testing"

	"github.com/sgago/col"
	"github.com/stretchr/testify/assert"
)

var kv1 col.KV[int] = col.KV[int]{Key: 1, Val: 1}
var kv2 col.KV[int] = col.KV[int]{Key: 2, Val: 2}
var kv3 col.KV[int] = col.KV[int]{Key: 3, Val: 3}
var kv4 col.KV[int] = col.KV[int]{Key: 4, Val: 4}
var kv5 col.KV[int] = col.KV[int]{Key: 5, Val: 5}

// TODO: Needs tests for when we add the the same value over and over again.

func TestNew_WithNoValues_CountIsZero(t *testing.T) {
	cap := 5

	monostack, _ := New[int](Decreasing, cap)

	assert.Zero(t, monostack.Count())
}

func TestNew_WithNoValues_CapacityIsCorrect(t *testing.T) {
	cap := 3

	monostack, _ := New[int](Decreasing, cap)

	assert.Equal(t, cap, monostack.Capacity())
}

func TestNew_WithDescValues_ElementsAreCorrect(t *testing.T) {
	values := []col.KV[int]{
		kv4,
		kv1,
		kv3,
		kv2,
	}

	monostack, _ := New(Decreasing, len(values), values...)

	assert.Equal(t, 3, monostack.Count())
	assert.Equal(t, kv4, monostack.elements[0])
	assert.Equal(t, kv3, monostack.elements[1])
	assert.Equal(t, kv2, monostack.elements[2])
}

func TestNew_WithAscValues_ElementsAreCorrect(t *testing.T) {
	values := []col.KV[int]{
		kv4,
		kv1,
		kv3,
		kv2,
	}

	monostack, _ := New(Increasing, len(values), values...)

	assert.Equal(t, 2, monostack.Count())
	assert.Equal(t, kv1, monostack.elements[0])
	assert.Equal(t, kv2, monostack.elements[1])
}

func TestNew_WithDescValues_PoppedValuesAreCorrect(t *testing.T) {
	values := []col.KV[int]{
		kv4,
		kv1,
		kv3,
		kv2,
	}

	_, popped := New(Decreasing, len(values), values...)

	assert.Equal(t, 1, len(popped))
	assert.Equal(t, kv1, popped[0])
}

func TestNew_WithAscValues_PoppedValuesAreCorrect(t *testing.T) {
	values := []col.KV[int]{
		kv4,
		kv1,
		kv3,
		kv2,
	}

	_, popped := New(Increasing, len(values), values...)

	assert.Equal(t, 2, len(popped))
	assert.Equal(t, kv4, popped[0])
	assert.Equal(t, kv3, popped[1])
}

func TestPush_WithDescValues_ElementsAreCorrect(t *testing.T) {
	cap := 5

	monostack, _ := New[int](Decreasing, cap)

	monostack.Push(kv3)
	monostack.Push(kv2)
	monostack.Push(kv4)
	monostack.Push(kv1)

	assert.Equal(t, 2, monostack.Count())
	assert.Equal(t, kv4, monostack.elements[0])
	assert.Equal(t, kv1, monostack.elements[1])
}

func TestPush_WithDescValues_PoppedValuesAreCorrect(t *testing.T) {
	cap := 5

	monostack, _ := New[int](Decreasing, cap)

	popped := make([]col.KV[int], 0)

	popped = append(popped, monostack.Push(kv3)...)
	popped = append(popped, monostack.Push(kv2)...)
	popped = append(popped, monostack.Push(kv4)...)
	popped = append(popped, monostack.Push(kv1)...)

	assert.Equal(t, 2, len(popped))
	assert.Equal(t, kv2, popped[0])
	assert.Equal(t, kv3, popped[1])
}

func TestPush_WithAscValues_ElementsAreCorrect(t *testing.T) {
	cap := 5

	monostack, _ := New[int](Increasing, cap)

	monostack.Push(kv2)
	monostack.Push(kv4)
	monostack.Push(kv1)
	monostack.Push(kv3)

	assert.Equal(t, 2, monostack.Count())
	assert.Equal(t, kv1, monostack.elements[0])
	assert.Equal(t, kv3, monostack.elements[1])
}

func TestPush_WithAscValues_PoppedValuesAreCorrect(t *testing.T) {
	cap := 5

	monostack, _ := New[int](Increasing, cap)

	popped := make([]col.KV[int], 0)

	popped = append(popped, monostack.Push(kv2)...)
	popped = append(popped, monostack.Push(kv4)...)
	popped = append(popped, monostack.Push(kv1)...)
	popped = append(popped, monostack.Push(kv3)...)

	assert.Equal(t, 2, len(popped))
	assert.Equal(t, kv4, popped[0])
	assert.Equal(t, kv2, popped[1])
}

func TestPop_WithValues_PoppedValueIsCorrect(t *testing.T) {
	cap := 3

	monostack, _ := New(Decreasing, cap, kv5, kv4, kv3)

	actual := monostack.Pop()

	assert.Equal(t, kv3, actual)
}

func TestPop_WithValues_ElementsAreCorrect(t *testing.T) {
	cap := 3

	monostack, _ := New(Decreasing, cap, kv5, kv4, kv3)

	monostack.Pop()

	assert.Equal(t, kv5, monostack.elements[0])
	assert.Equal(t, kv4, monostack.elements[1])
}

func TestPop_WithNilElements_Panics(t *testing.T) {
	stack := monostack[int]{}

	assert.Panics(t, func() { stack.Pop() })
}

func TestPop_WithZeroElements_Panics(t *testing.T) {
	stack := monostack[int]{elements: make([]col.KV[int], 0)}

	assert.Panics(t, func() { stack.Pop() })
}

func TestPeek_WithValues_PeekedValueIsCorrect(t *testing.T) {
	cap := 3

	monostack, _ := New(Decreasing, cap, kv5, kv4, kv3)

	actual := monostack.Peek()

	assert.Equal(t, kv3, actual)
}

func TestPeek_WithValues_ElementsAreCorrect(t *testing.T) {
	cap := 3

	monostack, _ := New(Decreasing, cap, kv5, kv4, kv3)

	monostack.Peek()

	assert.Equal(t, monostack.elements[0], kv5)
	assert.Equal(t, monostack.elements[1], kv4)
	assert.Equal(t, monostack.elements[2], kv3)
}

func TestPeek_WithNilElements_Panics(t *testing.T) {
	stack := monostack[int]{}

	assert.Panics(t, func() { stack.Peek() })
}

func TestPeek_WithZeroElements_Panics(t *testing.T) {
	stack := monostack[int]{elements: make([]col.KV[int], 0)}

	assert.Panics(t, func() { stack.Peek() })
}

func TestIsEmpty_WithElements_IsFalse(t *testing.T) {
	cap := 1

	monostack, _ := New(Decreasing, cap, kv1)

	assert.False(t, monostack.IsEmpty())
}

func TestIsEmpty_WithNoElements_IsTrue(t *testing.T) {
	cap := 1

	monostack, _ := New[int](Decreasing, cap)

	assert.True(t, monostack.IsEmpty())
}

func TestClear_WithElements_CountIsZero(t *testing.T) {
	cap := 3

	monostack, _ := New(Decreasing, cap, kv3, kv2, kv1)

	monostack.Clear()

	assert.Equal(t, 0, monostack.Count())
}

func TestClear_WithElements_CapacityIsSame(t *testing.T) {
	cap := 3

	monostack, _ := New(Decreasing, cap, kv3, kv2, kv1)

	monostack.Clear()

	assert.Equal(t, cap, monostack.Capacity())
}
