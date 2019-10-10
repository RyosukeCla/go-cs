package bstree

// BSTree is binary search tree
type BSTree struct {
	root *node
}

type node struct {
	value int
	left  *node
	right *node
}

// NewBSTree return BSTree
func NewBSTree() BSTree {
	return BSTree{}
}

// Insert inserts value
func (t *BSTree) Insert(value int) {
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

// Search searchs value
func (t *BSTree) Search(value int) bool {
	return search(t.root, value)
}

// Remove removes value
func (t *BSTree) Remove(value int) {
	if t.root == nil {
		return
	} else if t.root.value == value {
		t.root = swap(t.root)
	} else if t.root.value < value {
		remove(t.root.right, value)
	} else {
		remove(t.root.left, value)
	}
}

// Min returns min value
func (t *BSTree) Min() int {
	if t.root == nil {
		panic("error")
	}
	return mostLeft(t.root)
}

// Max returns max value
func (t *BSTree) Max() int {
	if t.root == nil {
		panic("error")
	}
	return mostRight(t.root)
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

func search(parent *node, value int) bool {
	if parent == nil {
		return false
	} else if parent.value == value {
		return true
	} else if parent.value < value {
		return search(parent.right, value)
	} else {
		return search(parent.left, value)
	}
}

func remove(parent *node, value int) {
	if parent.right == nil || parent.left == nil {
		return
	} else if parent.right.value == value {
		parent.right = swap(parent.right)
	} else if parent.left.value == value {
		parent.left = swap(parent.left)
	} else if parent.right.value > value {
		remove(parent.right, value)
	} else if parent.left.value < value {
		remove(parent.left, value)
	}
}

func swap(parent *node) *node {
	if parent.right == nil || parent.left == nil {
		return nil
	} else if parent.right != nil {
		return parent.right
	} else {
		return parent.left
	}
}

func mostLeft(parent *node) int {
	if parent.left == nil {
		return parent.value
	}
	return mostLeft(parent.left)
}

func mostRight(parent *node) int {
	if parent.right == nil {
		return parent.value
	}
	return mostRight(parent.right)
}
