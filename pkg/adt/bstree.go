package adt

type BSTree struct {
	root *node
}

type node struct {
	value float32
	left  *node
	right *node
}

func NewBSTree() BSTree {
	return BSTree{}
}

func (t *BSTree) Insert(value float32) {
	if t.root == nil {
		t.root = &node{
			value: value,
		}
	} else {
		insert(t.root, &node{
			value: value,
		})
	}
}

func insert(parent *node, child *node) {
	if parent.value < child.value {
		if parent.right == nil {
			parent.right = child
		} else {
			insert(parent.right, child)
		}
	} else {
		if parent.left == nil {
			parent.left = child
		} else {
			insert(parent.left, child)
		}
	}
}
