package treap

import (
	"math/rand"
)

// Treap is treap
type Treap[V interface{}] struct {
	root       *Node[V]
	comparator func(V, V) int
}

type Node[V interface{}] struct {
	priority int
	Value    V
	Left     *Node[V]
	Right    *Node[V]
}

// NewTreap return Treap
func NewTreap[V interface{}](comparator func(V, V) int) Treap[V] {
	return Treap[V]{
		comparator: comparator,
	}
}

// Insert inserts node
func (t *Treap[V]) Insert(value V) {
	t.root = insert(t.root, value, t.comparator)
}

// Search searches noe
func (t *Treap[V]) Search(value V) bool {
	return search(t.root, value, t.comparator)
}

// Remove removes node
func (t *Treap[V]) Remove(value V) {
	t.root = remove(t.root, value, t.comparator)
}

func (t *Treap[V]) Min() *V {
	if t.root == nil {
		return nil
	}
	res := mostLeft(t.root)
	return &res
}

func (t *Treap[V]) Max() *V {
	if t.root == nil {
		return nil
	}
	res := mostRight(t.root)
	return &res
}

func (t *Treap[V]) Walk(walker func(node *Node[V]) int) {
	walk(t.root, walker)
}

func walk[V interface{}](n *Node[V], walker func(n *Node[V]) int) {
	if n == nil {
		return
	}
	compared := walker(n)
	if compared == 0 {
		return
	}
	if compared < 0 {
		walk(n.Left, walker)
	} else {
		walk(n.Right, walker)
	}
}

func insert[V interface{}](root *Node[V], value V, comparator func(V, V) int) *Node[V] {
	if root == nil {
		return createNode(value)
	}
	if comparator(root.Value, value) >= 0 {
		root.Left = insert(root.Left, value, comparator)
		if root.Left.priority > root.priority {
			root = rightRotate(root)
		}
	} else {
		root.Right = insert(root.Right, value, comparator)
		if root.Right.priority > root.priority {
			root = leftRotate(root)
		}
	}
	return root
}

func rightRotate[V interface{}](root *Node[V]) *Node[V] {
	newRoot := root.Left
	tree := newRoot.Right
	// rotation
	newRoot.Right = root
	root.Left = tree
	return newRoot
}

func leftRotate[V interface{}](root *Node[V]) *Node[V] {
	newRoot := root.Right
	tree := newRoot.Left
	// rotation
	newRoot.Left = root
	root.Right = tree
	return newRoot
}

func search[V interface{}](parent *Node[V], value V, comparator func(V, V) int) bool {
	if parent == nil {
		return false
	}
	compared := comparator(parent.Value, value)
	if compared == 0 {
		return true
	} else if compared < 0 {
		return search(parent.Right, value, comparator)
	} else {
		return search(parent.Left, value, comparator)
	}
}

func remove[V interface{}](parent *Node[V], value V, comparator func(V, V) int) *Node[V] {
	if parent == nil {
		return nil
	}

	compared := comparator(parent.Value, value)
	if compared > 0 {
		parent.Left = remove(parent.Left, value, comparator)
		return parent
	} else if compared < 0 {
		parent.Right = remove(parent.Right, value, comparator)
		return parent
	} else if parent.Left == nil {
		return parent.Right
	} else if parent.Right == nil {
		return parent.Left
	} else if parent.Left.priority < parent.Right.priority {
		newRoot := leftRotate(parent)
		newRoot.Left = remove(newRoot.Left, value, comparator)
		return newRoot
	} else {
		newRoot := rightRotate(parent)
		newRoot.Right = remove(newRoot.Right, value, comparator)
		return newRoot
	}
}

func mostLeft[V interface{}](parent *Node[V]) V {
	if parent.Left == nil {
		return parent.Value
	}
	return mostLeft(parent.Left)
}

func mostRight[V interface{}](parent *Node[V]) V {
	if parent.Right == nil {
		return parent.Value
	}
	return mostRight(parent.Right)
}

func createNode[V interface{}](value V) *Node[V] {
	return &Node[V]{
		Value:    value,
		priority: rand.Intn(1000000),
	}
}
