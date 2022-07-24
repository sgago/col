package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf_ValueDoesNotExist_ReturnedIndexNegative1(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)

	actual, _ := IndexOf(s, 6)

	assert.Equal(t, -1, actual)
}

func TestIndexOf_ValueExists_ReturnsIndex(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	values := make([]int, 0, count)

	for i := 0; i < count; i++ {
		values = append(values, i)
	}

	i, _ := IndexOf(values, count-1)

	assert.Equal(t, count-1, i)
}

func TestIndexOf_ValueExists_ReturnsNilError(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	values := make([]int, 0, count)

	for i := 0; i < count; i++ {
		values = append(values, i)
	}

	_, e := IndexOf(values, count-1)

	assert.Nil(t, e)
}

func TestIndexOf_ValueDoesNotExist_ReturnsNegativeOneIndex(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	values := make([]int, 0, count)

	for i := 0; i < count; i++ {
		values = append(values, i)
	}

	i, _ := IndexOf(values, count+1)

	assert.Equal(t, i, -1)
}

func TestIndexOf_ValueDoesNotExist_ReturnsError(t *testing.T) {
	count := 3 * DefaultMaxSearchLength

	values := make([]int, 0, count)

	for i := 0; i < count; i++ {
		values = append(values, i)
	}

	_, e := IndexOf(values, count+1)

	assert.NotNil(t, e)
}
