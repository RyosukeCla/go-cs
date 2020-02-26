package linkedlist

type Node struct {
	prev, next *Node
	Value      interface{}
}

func newNode(value interface{}) *Node {
	return &Node{
		Value: value,
	}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	first, last *Node
	len         int
}

func New() List {
	return List{
		len: 0,
	}
}

func (l *List) Len() int {
	return l.len
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}

func (l *List) InsertAfter(node *Node, value interface{}) *Node {
	// node <- newNode -> node.next
	newNode := newNode(value)
	newNode.prev = node
	newNode.next = node.next

	// newNode <- node.next
	if node.next != nil {
		node.next.prev = newNode
	}

	// node -> newNode
	node.next = newNode

	if l.last == node {
		l.last = newNode
	}

	l.len++

	return newNode
}

func (l *List) InsertBefore(node *Node, value interface{}) *Node {
	// node.prev <- newNode -> node
	newNode := newNode(value)
	newNode.prev = node.prev
	newNode.next = node

	// node.prev -> newNode <- node
	if node.prev != nil {
		node.prev.next = newNode
	}

	// newNode <- node
	node.prev = newNode

	if l.first == node {
		l.first = newNode
	}

	l.len++

	return newNode
}

func (l *List) InsertBiginning(value interface{}) *Node {
	if l.first == nil {
		node := newNode(value)
		l.first = node
		l.last = node
		l.len++
		return node
	}

	return l.InsertBefore(l.first, value)
}

func (l *List) InsertEnd(value interface{}) *Node {
	if l.last == nil {
		return l.InsertBiginning(value)
	}
	return l.InsertAfter(l.last, value)
}

func (l *List) MoveToBeginning(node *Node) {
	l.Remove(node)

	// node <-> fisrt
	node.next = l.first
	l.first.prev = node
	l.first = node
}

func (l *List) MoveToEnd(node *Node) {
	l.Remove(node)

	// last <-> node
	node.prev = l.last
	l.last.next = node
	l.last = node
}

func (l *List) Remove(node *Node) {
	if l.first == node {
		l.first = node.next
	}
	if l.last == node {
		l.last = node.prev
	}

	// node.prev <-> node.next
	next := node.next
	prev := node.prev
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}

	node.next = nil
	node.prev = nil

	l.len--
}

func (l *List) Swap(a, b *Node) {
	if l.first == a {
		l.first = b
	}
	if l.first == b {
		l.first = a
	}
	if l.last == a {
		l.last = b
	}
	if l.last == b {
		l.last = a
	}

	// a.prev <-> b <-> a.next ... b.prev <-> a <-> b.next
	aPrev := a.prev
	aNext := a.next
	bPrev := b.prev
	bNext := b.next

	// a.prev <-> b
	b.prev = nil
	if aPrev != nil {
		aPrev.next = b
		b.prev = aPrev
	}

	// b <-> a.next
	b.next = nil
	if aNext != nil {
		aNext.prev = b
		b.next = aNext
	}

	// b.prev <-> a
	a.prev = nil
	if bPrev != nil {
		bPrev.next = a
		a.prev = bPrev
	}

	// a <-> b.next
	a.next = nil
	if bNext != nil {
		bNext.prev = a
		a.next = bNext
	}
}
