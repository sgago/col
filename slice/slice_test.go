package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFirst(t *testing.T) {
	s := make([]int, 0, 2)
	s = append(s, 111, 222)

	s = RemoveFirst(s)

	assert.Equal(t, 1, len(s))
	assert.Equal(t, s[0], 222)
}

func TestRemoveLast(t *testing.T) {
	s := make([]int, 0, 2)
	s = append(s, 111, 222)

	s = RemoveLast(s)

	assert.Equal(t, 1, len(s))
	assert.Equal(t, s[0], 111)
}

func TestSwap(t *testing.T) {
	s := make([]int, 0, 2)
	s = append(s, 111, 222, 333)

	s = Swap(s, 0, 2)

	assert.Equal(t, 3, len(s))
	assert.Equal(t, s[0], 333)
	assert.Equal(t, s[1], 222)
	assert.Equal(t, s[2], 111)
}

func TestContains_ValueDoesNotExist_ReturnsFalse(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	actual := Contains(s, 6)

	assert.False(t, actual)
}

func TestContains_ValueExists_ReturnsTrue(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	actual := Contains(s, 5)

	assert.True(t, actual)
}

func TestClearLengthIsZero(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	clear := Clear(s)

	assert.Equal(t, len(clear), 0)
}

func TestClearCapacityRemains(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	clear := Clear(s)

	assert.Equal(t, cap(clear), 5)
}

func TestAny_WithPredicateReturningTrue_ReturnsTrue(t *testing.T) {
	values := make([]bool, 3)

	values[0] = false
	values[1] = false
	values[2] = true

	predicate := func(index int, value bool) bool {
		return value
	}

	actual := Any(values, predicate)

	assert.True(t, actual)
}

func TestAny_WithPredicateReturningFalse_ReturnsFalse(t *testing.T) {
	values := make([]bool, 3)

	values[0] = false
	values[1] = false
	values[2] = false

	predicate := func(index int, value bool) bool {
		return value
	}

	actual := Any(values, predicate)

	assert.False(t, actual)
}

func TestAny_WithNilPredicateAndSliceHasElements_ReturnsTrue(t *testing.T) {
	values := make([]bool, 3)

	values[0] = false
	values[1] = false
	values[2] = false

	actual := Any(values, nil)

	assert.True(t, actual)
}

func TestAny_WithNilPredicateAndSliceHasNoElements_ReturnsTrue(t *testing.T) {
	values := make([]bool, 0)

	actual := Any(values, nil)

	assert.False(t, actual)
}

func TestAll_WithPredicateAlwaysReturningTrue_ReturnsTrue(t *testing.T) {
	values := make([]bool, 3)

	values[0] = true
	values[1] = true
	values[2] = true

	predicate := func(index int, value bool) bool {
		return value
	}

	actual := All(values, predicate)

	assert.True(t, actual)
}

func TestAll_WithPredicateReturningFalse_ReturnsFalse(t *testing.T) {
	values := make([]bool, 3)

	values[0] = true
	values[1] = true
	values[2] = false

	predicate := func(index int, value bool) bool {
		return value
	}

	actual := All(values, predicate)

	assert.False(t, actual)
}

func TestAll_WithNilPredicate_Panics(t *testing.T) {
	values := make([]bool, 3)

	values[0] = false
	values[1] = false
	values[2] = false

	assert.Panics(t, func() { All(values, nil) })
}
