package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueueAndDequeue(t *testing.T) {
	q := New(3, 1, 2, 3)

	a := q.Dequeue()
	b := q.Dequeue()
	c := q.Dequeue()

	assert.Equal(t, a, 1)
	assert.Equal(t, b, 2)
	assert.Equal(t, c, 3)
}
