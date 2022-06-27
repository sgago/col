package binarytree

import (
	"testing"

	"github.com/sgago/col"
	"github.com/stretchr/testify/assert"
)

var pv1 col.PV[int] = col.PV[int]{Priority: 1, Val: 1}
var pv2 col.PV[int] = col.PV[int]{Priority: 2, Val: 2}
var pv3 col.PV[int] = col.PV[int]{Priority: 3, Val: 3}
var pv4 col.PV[int] = col.PV[int]{Priority: 4, Val: 4}

func TestInsert_WithPriorityLessThanCurrentNode_CreatesLeftNode(t *testing.T) {
	bt := New(pv2)

	bt.Insert(pv1)

	assert.Equal(t, pv1.Priority, bt.left.Priority)
}

func TestInsert_WithPriorityGreaterThanCurrentNode_CreatesRightNode(t *testing.T) {
	bt := New(pv2)

	bt.Insert(pv3)

	assert.Equal(t, pv3.Priority, bt.right.Priority)
}

func TestFind_WithPriorityInTree_NodeIsFound(t *testing.T) {
	bt := New(pv2)

	bt.Insert(pv1)
	bt.Insert(pv3)
	bt.Insert(pv4)

	node, _ := bt.Find(4)

	assert.Equal(t, pv4.Priority, node.Priority)
}

func TestFind_WithPriorityInTree_ErrorIsNil(t *testing.T) {
	bt := New(pv2)

	bt.Insert(pv1)
	bt.Insert(pv3)
	bt.Insert(pv4)

	_, e := bt.Find(4)

	assert.Nil(t, e)
}

func TestFind_WithPriorityNotInTree_ReturnsNilNode(t *testing.T) {
	bt := New(pv2)

	bt.Insert(pv1)
	bt.Insert(pv3)
	bt.Insert(pv4)

	node, _ := bt.Find(7)

	var expected int

	assert.Equal(t, expected, node.Priority)
}

func TestFind_WithPriorityNotInTree_ReturnsError(t *testing.T) {
	bt := New(pv2)

	bt.Insert(pv1)
	bt.Insert(pv3)
	bt.Insert(pv4)

	_, notFoundError := bt.Find(7)

	assert.NotNil(t, notFoundError)
}

func TestRemove(t *testing.T) {
	bt := New(pv2)

	bt.Insert(pv1)
	bt.Insert(pv3)
	bt.Insert(pv4)

	bt.Remove(7)
}
