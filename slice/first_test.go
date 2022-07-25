package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst_WithNilPredicate_ReturnsFirstElement(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, first, _ := First(s, nil)

	assert.Equal(t, first, 1)
}

func TestFirst_WithPredicateReturningTrue_ReturnsFirstElement(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, first, _ := First(s, func(i int, v int) bool {
		return v == 3
	})

	assert.Equal(t, first, 3)
}

func TestFirst_WithPredicateReturningTrue_ReturnsNilError(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, _, e := First(s, func(i int, v int) bool {
		return v == 3
	})

	assert.Nil(t, e)
}

func TestFirst_WithPredicateReturningFalse_ReturnsDefault(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, first, _ := First(s, func(i int, v int) bool {
		return v == 123
	})

	assert.Equal(t, first, 0)
}

func TestFirst_WithPredicateReturningFalse_ReturnsError(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, _, e := First(s, func(i int, v int) bool {
		return v == 123
	})

	assert.NotNil(t, e)
}

func TestFirst_With0LengthSlice_Panics(t *testing.T) {
	s := make([]int, 0)

	assert.Panics(t, func() { First(s, nil) })
}

func TestFirst_WithGoRoutines_PredicateReturningTrue_ReturnsFirstIndex(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, 0, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	index, _, _ := First(s, func(i int, v int) bool {
		return i == count-1
	})

	assert.Equal(t, count-1, index)
}

func TestFirst_WithGoRoutines_PredicateReturningTrue_ReturnsFirstValue(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, 0, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	_, value, _ := First(s, func(i int, v int) bool {
		return i == count-1
	})

	assert.Equal(t, count-1, value)
}

func TestFirst_WithGoRoutines_PredicateReturningTrue_ReturnsNilError(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, 0, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	_, _, e := First(s, func(i int, v int) bool {
		return i == count-1
	})

	assert.Nil(t, e)
}

func TestFirst_WithGoRoutines_PredicateReturningFalse_ReturnsNotFoundIndex(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, 0, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	index, _, _ := First(s, func(i int, v int) bool {
		return false
	})

	assert.Equal(t, NotFound, index)
}

func TestFirst_WithGoRoutines_PredicateReturningFalse_ReturnsDefaultValue(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, 0, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	_, value, _ := First(s, func(i int, v int) bool {
		return false
	})

	assert.Equal(t, 0, value)
}

func TestFirst_WithGoRoutines_PredicateReturningFalse_ReturnsError(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, 0, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	_, _, e := First(s, func(i int, v int) bool {
		return false
	})

	assert.NotNil(t, e)
}
