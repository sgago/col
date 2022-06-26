package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst_WithNilPredicate_ReturnsFirstElement(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	first, _ := First(s, nil)

	assert.Equal(t, first, 1)
}

func TestFirst_WithPredicateReturningTrue_ReturnsFirstElement(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	first, _ := First(s, func(i int, v int) bool {
		return v == 3
	})

	assert.Equal(t, first, 3)
}

func TestFirst_WithPredicateReturningTrue_ReturnsNilError(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, e := First(s, func(i int, v int) bool {
		return v == 3
	})

	assert.Nil(t, e)
}

func TestFirst_WithPredicateReturningFalse_ReturnsDefault(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	first, _ := First(s, func(i int, v int) bool {
		return v == 123
	})

	assert.Equal(t, first, 0)
}

func TestFirst_WithPredicateReturningFalse_ReturnsError(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, e := First(s, func(i int, v int) bool {
		return v == 123
	})

	assert.NotNil(t, e)
}

func TestFirst_With0LengthSlice_Panics(t *testing.T) {
	s := make([]int, 0)

	assert.Panics(t, func() { First(s, nil) })
}

func TestLast_WithNilPredicate_ReturnsFirstElement(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	last, _ := Last(s, nil)

	assert.Equal(t, last, 5)
}

func TestLast_WithPredicateReturningTrue_ReturnsFirstElement(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	last, _ := Last(s, func(i int, v int) bool {
		return v == 3
	})

	assert.Equal(t, last, 3)
}

func TestLast_WithPredicateReturningTrue_ReturnsNilError(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, e := Last(s, func(i int, v int) bool {
		return v == 3
	})

	assert.Nil(t, e)
}

func TestLast_WithPredicateReturningFalse_ReturnsDefault(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	last, _ := Last(s, func(i int, v int) bool {
		return v == 123
	})

	assert.Equal(t, last, 0)
}

func TestLast_WithPredicateReturningFalse_ReturnsError(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, e := Last(s, func(i int, v int) bool {
		return v == 123
	})

	assert.NotNil(t, e)
}

func TestLast_With0LengthSlice_Panics(t *testing.T) {
	s := make([]int, 0)

	assert.Panics(t, func() { Last(s, nil) })
}

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

func TestIndexOf_ValueDoesNotExist_ReturnedIndexNegative1(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	actual, _ := IndexOf(s, 6)

	assert.Equal(t, -1, actual)
}

func TestIndexOf_ValueExists_ReturnsIndex(t *testing.T) {
	count := 30_000

	values := make([]int, 0, count)

	for i := 0; i < count; i++ {
		values = append(values, i)
	}

	i, _ := IndexOf(values, count-1)

	assert.Equal(t, i, count-1)
}

func TestIndexOf_ValueExists_ReturnsNilError(t *testing.T) {
	count := 30_000

	values := make([]int, 0, count)

	for i := 0; i < count; i++ {
		values = append(values, i)
	}

	_, e := IndexOf(values, count-1)

	assert.Nil(t, e)
}

func TestIndexOf_ValueDoesNotExist_ReturnsNegativeOneIndex(t *testing.T) {
	count := 30_000

	values := make([]int, 0, count)

	for i := 0; i < count; i++ {
		values = append(values, i)
	}

	i, _ := IndexOf(values, count+1)

	assert.Equal(t, i, -1)
}

func TestIndexOf_ValueDoesNotExist_ReturnsError(t *testing.T) {
	count := 30_000

	values := make([]int, 0, count)

	for i := 0; i < count; i++ {
		values = append(values, i)
	}

	_, e := IndexOf(values, count+1)

	assert.NotNil(t, e)
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
