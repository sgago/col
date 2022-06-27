package binarytree

import (
	"github.com/sgago/col"
	"github.com/sgago/col/err"
)

type node[T any] struct {
	col.PV[T]
	left  *node[T]
	right *node[T]
}

func New[T any](pv col.PV[T]) *node[T] {
	return &node[T]{PV: pv}
}

func (n *node[T]) Insert(pv col.PV[T]) {

	if pv.Priority < n.Priority {
		if n.left == nil {
			n.left = &node[T]{PV: pv}
		} else {
			n.left.Insert(pv)
		}
	} else {
		if n.right == nil {
			n.right = &node[T]{PV: pv}
		} else {
			n.right.Insert(pv)
		}
	}
}

func (n *node[T]) Find(priority int) (node[T], error) {
	if n.Priority == priority {
		return *n, nil
	}

	if n.isLeaf() {
		return node[T]{}, &err.KeyNotFound{Key: priority}
	}

	if n.Priority > priority {
		return n.left.Find(priority)
	} else {
		return n.right.Find(priority)
	}
}

func (root *node[T]) Remove(priority int) error {
	if root.Priority == priority {
		// This is a special case where we're deleting the root node
		if root.left != nil {
			root.left.right = root.right
			*root = *root.left
		} else if root.right != nil {
			*root = *root.right
		} else {
			root = nil
		}

		return nil
	}

	return root.remove(priority, root)
}

func (n *node[T]) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *node[T]) remove(priority int, parent *node[T]) error {
	if n.Priority == priority {
		n.deleteNode(parent)
		return nil
	}

	if n.isLeaf() && n.Priority != priority {
		return &err.KeyNotFound{Key: priority}
	}

	if n.Priority > priority {
		return n.left.remove(priority, n)
	} else {
		return n.right.remove(priority, n)
	}
}

func (n *node[T]) deleteNode(parent *node[T]) {
	isLeftChild := parent.left == n

	if n.left == nil && n.right != nil {
		if isLeftChild {
			parent.left = n.right
		} else {
			parent.right = n.right
		}
	} else if n.left != nil && n.right == nil {
		if isLeftChild {
			parent.left = n.left
		} else {
			parent.right = n.left
		}
	} else {
		// TODO: these all need tests
		if isLeftChild {
			parent.left = n.left

			if n.left != nil {
				n.left.right = n.right
			}
		} else {
			parent.right = n.left

			if n.left != nil {
				n.left.right = n.right
			}
		}
	}

	n = nil
}
