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

// NewTreap return Treap
func NewTreap() Treap {
	return Treap{}
}

// Insert inserts node
func (t *Treap) Insert(key int) {
	t.root = insert(t.root, key)
}

// Search searches noe
func (t *Treap) Search(key int) bool {
	return search(t.root, key)
}

// Remove removes node
func (t *Treap) Remove(key int) {
	remove(t.root, key)
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

func search(parent *node, key int) bool {
	if parent == nil {
		return false
	} else if parent.key == key {
		return true
	} else if parent.key < key {
		return search(parent.right, key)
	} else {
		return search(parent.left, key)
	}
}

func remove(parent *node, key int) *node {
	if parent == nil {
		return nil
	}

	if key < parent.key {
		parent.left = remove(parent.left, key)
	} else if parent.key < key {
		parent.right = remove(parent.right, key)
	} else if parent.left == nil {
		newParent := parent.right
		parent.right = nil
		return newParent
	} else if parent.right == nil {
		newParent := parent.left
		parent.left = nil
		return newParent
	} else if parent.left.value < parent.right.value {
		newParent := leftRotate(parent)
		newParent.left = remove(newParent.left, key)
		return newParent
	} else {
		newParent := rightRotate(parent)
		newParent.right = remove(newParent.right, key)
		return newParent
	}

	return parent
}

func createNode(key int) *node {
	return &node{
		key:   key,
		value: rand.Intn(1000000),
	}
}
