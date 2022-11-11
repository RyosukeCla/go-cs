package bstree

// BSTree is binary search tree
type BSTree[V interface{}] struct {
	root       *Node[V]
	comparator func(V, V) int
}

type Node[V interface{}] struct {
	Value V
	Left  *Node[V]
	Right *Node[V]
}

func NumberComparator[V int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](left V, right V) int {
	if left == right {
		return 0
	} else if left < right {
		return -1
	}
	return 1
}

// NewBSTree return BSTree
func NewBSTree[V interface{}](comparator func(V, V) int) BSTree[V] {
	return BSTree[V]{
		comparator: comparator,
	}
}

// Insert inserts value
func (t *BSTree[V]) Insert(value V) {
	if t.root == nil {
		t.root = &Node[V]{
			Value: value,
		}
	} else {
		t.insert(t.root, &Node[V]{
			Value: value,
		})
	}
}

// Search searchs value
func (t *BSTree[V]) Search(value V) bool {
	return t.search(t.root, value)
}

func (t *BSTree[V]) Walk(walker func(node *Node[V]) int) {
	walk(t.root, walker)
}

// Remove removes value
func (t *BSTree[V]) Remove(value V) {
	if t.root == nil {
		return
	} else if t.comparator(t.root.Value, value) == 0 {
		t.root = t.swap(t.root)
	} else if t.comparator(t.root.Value, value) == -1 {
		t.remove(t.root.Right, value)
	} else {
		t.remove(t.root.Left, value)
	}
}

// Min returns min value
func (t *BSTree[V]) Min() V {
	if t.root == nil {
		panic("error")
	}
	return t.mostLeft(t.root)
}

// Max returns max value
func (t *BSTree[V]) Max() V {
	if t.root == nil {
		panic("error")
	}
	return t.mostRight(t.root)
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

func (t *BSTree[V]) insert(parent *Node[V], child *Node[V]) {
	if t.comparator(parent.Value, child.Value) == -1 {
		if parent.Right == nil {
			parent.Right = child
		} else {
			t.insert(parent.Right, child)
		}
	} else {
		if parent.Left == nil {
			parent.Left = child
		} else {
			t.insert(parent.Left, child)
		}
	}
}

func (t *BSTree[V]) search(parent *Node[V], value V) bool {
	if parent == nil {
		return false
	} else if t.comparator(parent.Value, value) == 0 {
		return true
	} else if t.comparator(parent.Value, value) == -1 {
		return t.search(parent.Right, value)
	} else {
		return t.search(parent.Left, value)
	}
}

func (t *BSTree[V]) remove(parent *Node[V], value V) {
	if parent.Right == nil || parent.Left == nil {
		return
	} else if t.comparator(parent.Right.Value, value) == 0 {
		parent.Right = t.swap(parent.Right)
	} else if t.comparator(parent.Left.Value, value) == 0 {
		parent.Left = t.swap(parent.Left)
	} else if t.comparator(parent.Right.Value, value) == 1 {
		t.remove(parent.Right, value)
	} else if t.comparator(parent.Left.Value, value) == -1 {
		t.remove(parent.Left, value)
	}
}

func (t *BSTree[V]) swap(parent *Node[V]) *Node[V] {
	if parent.Right == nil || parent.Left == nil {
		return nil
	} else if parent.Right != nil {
		return parent.Right
	} else {
		return parent.Left
	}
}

func (t *BSTree[V]) mostLeft(parent *Node[V]) V {
	if parent.Left == nil {
		return parent.Value
	}
	return t.mostLeft(parent.Left)
}

func (t *BSTree[V]) mostRight(parent *Node[V]) V {
	if parent.Right == nil {
		return parent.Value
	}
	return t.mostRight(parent.Right)
}
