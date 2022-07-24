package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLast_WithNilPredicate_ReturnsFirstElement(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, last, _ := Last(s, nil)

	assert.Equal(t, last, 5)
}

func TestLast_WithPredicateReturningTrue_ReturnsLastElement(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, last, _ := Last(s, func(i int, v int) bool {
		return v == 3
	})

	assert.Equal(t, last, 3)
}

func TestLast_WithPredicateReturningTrue_ReturnsNilError(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, _, e := Last(s, func(i int, v int) bool {
		return v == 3
	})

	assert.Nil(t, e)
}

func TestLast_WithPredicateReturningFalse_ReturnsDefault(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, last, _ := Last(s, func(i int, v int) bool {
		return v == 123
	})

	assert.Equal(t, last, 0)
}

func TestLast_WithPredicateReturningFalse_ReturnsError(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	_, _, e := Last(s, func(i int, v int) bool {
		return v == 123
	})

	assert.NotNil(t, e)
}

func TestLast_With0LengthSlice_Panics(t *testing.T) {
	s := make([]int, 0)

	assert.Panics(t, func() { Last(s, nil) })
}

func TestLast_WithGoRoutines_PredicateReturningTrue_ReturnsFirstIndex(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	index, _, _ := Last(s, func(i int, v int) bool {
		return i == 0
	})

	assert.Equal(t, 0, index)
}

func TestLast_WithGoRoutines_PredicateReturningTrue_ReturnsFirstValue(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	_, value, _ := Last(s, func(i int, v int) bool {
		return i == 0
	})

	assert.Equal(t, 0, value)
}

func TestLast_WithGoRoutines_PredicateReturningTrue_ReturnsNilError(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	_, _, e := Last(s, func(i int, v int) bool {
		return i == 0
	})

	assert.Nil(t, e)
}

func TestLast_WithGoRoutines_PredicateReturningFalse_ReturnsNotFoundIndex(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	index, _, _ := Last(s, func(i int, v int) bool {
		return false
	})

	assert.Equal(t, NotFound, index)
}

func TestLast_WithGoRoutines_PredicateReturningFalse_ReturnsDefaultValue(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	_, value, _ := Last(s, func(i int, v int) bool {
		return false
	})

	assert.Equal(t, 0, value)
}

func TestLast_WithGoRoutines_PredicateReturningFalse_ReturnsError(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	s := make([]int, count)

	for i := 0; i < count; i++ {
		s = append(s, i)
	}

	_, _, e := Last(s, func(i int, v int) bool {
		return false
	})

	assert.NotNil(t, e)
}
