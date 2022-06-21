package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_WithNoValues_CountIsZero(t *testing.T) {
	cap := 3

	stack := New[int](cap)

	assert.Zero(t, stack.Count())
}

func TestNew_WithNoValues_CapacityIsCorrect(t *testing.T) {
	cap := 3

	stack := New[int](cap)

	assert.Equal(t, cap, stack.Capacity())
}

func TestNew_WithValues_ItemsAreCorrect(t *testing.T) {
	cap := 5

	values := make([]int, 0, cap)

	for i := 0; i < cap; i++ {
		values = append(values, i)
	}

	stack := New(cap, values...)

	for i := 0; i < len(values); i++ {
		assert.Equal(t, stack.items[i], values[i])
	}
}

func TestPush_WithValues_ItemsAreCorrect(t *testing.T) {
	cap := 5

	stack := New[int](cap)

	values := make([]int, 0, cap)

	for i := 0; i < cap; i++ {
		stack.Push(i)
	}

	for i := 0; i < len(values); i++ {
		assert.Equal(t, stack.items[i], values[i])
	}
}

func TestPop_WithValues_PoppedValueIsCorrect(t *testing.T) {
	cap := 5

	values := make([]int, 0, cap)

	for i := 0; i < cap; i++ {
		values = append(values, i)
	}

	stack := New(cap, values...)

	actual := stack.Pop()

	assert.Equal(t, values[len(values)-1], actual)
}

func TestPop_WithValues_ItemsAreCorrect(t *testing.T) {
	cap := 5

	values := make([]int, 0, cap)

	for i := 0; i < cap; i++ {
		values = append(values, i)
	}

	stack := New(cap, values...)

	stack.Pop()

	for i := 0; i < len(values)-1; i++ {
		assert.Equal(t, stack.items[i], values[i])
	}
}

func TestPeek_WithValues_PeekedValueIsCorrect(t *testing.T) {
	cap := 5

	values := make([]int, 0, cap)

	for i := 0; i < cap; i++ {
		values = append(values, i)
	}

	stack := New(cap, values...)

	actual := stack.Peek()

	assert.Equal(t, values[len(values)-1], actual)
}

func TestPeek_WithValues_ItemsAreCorrect(t *testing.T) {
	cap := 5

	values := make([]int, 0, cap)

	for i := 0; i < cap; i++ {
		values = append(values, i)
	}

	stack := New(cap, values...)

	stack.Peek()

	for i := 0; i < len(values); i++ {
		assert.Equal(t, stack.items[i], values[i])
	}
}

func TestIsEmpty_WithElements_IsFalse(t *testing.T) {
	cap := 1

	stack := New(cap, 1)

	assert.False(t, stack.IsEmpty())
}

func TestIsEmpty_WithNoElements_IsTrue(t *testing.T) {
	cap := 1

	stack := New[int](cap)

	assert.True(t, stack.IsEmpty())
}

func TestClear_WithElements_CountIsZero(t *testing.T) {
	cap := 3

	stack := New(cap, 1, 2, 3)

	stack.Clear()

	assert.Equal(t, 0, stack.Count())
}

func TestClear_WithElements_CapacityIsSame(t *testing.T) {
	cap := 3

	stack := New(cap, 1, 2, 3)

	stack.Clear()

	assert.Equal(t, cap, stack.Capacity())
}
