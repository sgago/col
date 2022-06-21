package binarytree

import "testing"

func TestInsert_WithKeyLessThanCurrentNode_CreatesLeftNode(t *testing.T) {
	bt := New(2, 2)

	bt.Insert(1, 1)

	if bt.left.Key != 1 {
		t.Errorf("The left node should have a key of 1")
	}
}

func TestInsert_WithKeyGreaterThanCurrentNode_CreatesRightNode(t *testing.T) {
	bt := New(2, 2)

	bt.Insert(3, 3)

	if bt.right.Key != 3 {
		t.Errorf("The right node should have a key of 3")
	}
}

func TestFind_WithKeyInTree_NodeIsFound(t *testing.T) {

	bt := New(2, 2)

	bt.Insert(1, 1)
	bt.Insert(3, 3)
	bt.Insert(4, 4)

	node, notFoundError := bt.Find(4)

	if notFoundError != nil {
		t.Errorf("Expected to find %d, but it was not found", node.Key)
	}
}

func TestFind_WithKeyNotInTree_ReturnsError(t *testing.T) {

	bt := New(2, 2)

	bt.Insert(1, 1)
	bt.Insert(3, 3)
	bt.Insert(4, 4)

	_, notFoundError := bt.Find(7)

	if notFoundError == nil {
		t.Error("Expected a not found error but none was returned")
	}
}

func TestRemove(t *testing.T) {
	bt := New(2, 2)

	bt.Insert(1, 1)
	bt.Insert(3, 3)
	bt.Insert(4, 4)

	bt.Remove(7)
}
