package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirst(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	first := First(s)

	assert.Equal(t, first, 1)
}

func TestLast(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	last := Last(s)

	assert.Equal(t, last, 5)
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
