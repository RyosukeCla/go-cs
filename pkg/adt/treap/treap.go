package treap

import "math/rand"

// Treap is treap
type Treap struct {
	root *node
}

type node struct {
	key   int
	value int
	left  *node
	right *node
}

func NewTreap() Treap {
	return Treap{}
}

func (t *Treap) Insert(key int) {
	t.root = insert(t.root, key)
}

func insert(root *node, key int) *node {
	if root == nil {
		return createNode(key)
	}

	if key < root.key {
		root.left = insert(root.left, key)

		if root.left.value > root.value {
			root = rightRotate(root)
		}
	} else {
		root.right = insert(root.right, key)

		if root.right.value > root.value {
			root = leftRotate(root)
		}
	}

	return root
}

func rightRotate(root *node) *node {
	newRoot := root.left
	tree := newRoot.right
	newRoot.right = root
	root.left = tree
	return newRoot
}

func leftRotate(root *node) *node {
	newRoot := root.right
	tree := newRoot.left
	newRoot.left = root
	root.right = tree
	return newRoot
}

func createNode(key int) *node {
	return &node{
		key:   key,
		value: rand.Intn(1000000),
	}
}
