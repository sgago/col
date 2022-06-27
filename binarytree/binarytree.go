package binarytree

import (
	"github.com/sgago/col"
	"github.com/sgago/col/err"
)

type node[T any] struct {
	col.KV[int, T]
	left  *node[T]
	right *node[T]
}

func (n *node[T]) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func New[T any](key int, value T) *node[T] {
	return &node[T]{KV: col.KV[int, T]{Key: key, Val: value}}
}

func (n *node[T]) Insert(key int, value T) {

	if key < n.Key {
		if n.left == nil {
			n.left = &node[T]{KV: col.KV[int, T]{Key: key, Val: value}}
		} else {
			n.left.Insert(key, value)
		}
	} else {
		if n.right == nil {
			n.right = &node[T]{KV: col.KV[int, T]{Key: key, Val: value}}
		} else {
			n.right.Insert(key, value)
		}
	}
}

func (n *node[T]) Find(key int) (node[T], error) {
	if n.Key == key {
		return *n, nil
	}

	if n.isLeaf() {
		return node[T]{}, &err.KeyNotFound{Key: key}
	}

	if n.Key > key {
		return n.left.Find(key)
	} else {
		return n.right.Find(key)
	}
}

func (root *node[T]) Remove(key int) error {
	if root.Key == key {
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

	return root.remove(key, root)
}

func (n *node[T]) remove(key int, parent *node[T]) error {
	if n.Key == key {
		n.deleteNode(parent)
		return nil
	}

	if n.isLeaf() && n.Key != key {
		return &err.KeyNotFound{Key: key}
	}

	if n.Key > key {
		return n.left.remove(key, n)
	} else {
		return n.right.remove(key, n)
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
