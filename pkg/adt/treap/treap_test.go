package treap

import (
	"fmt"
	"testing"
)

func TestTreap(t *testing.T) {
	treap := NewTreap()
	treap.Insert(1)
	treap.Insert(2)
	treap.Insert(3)
	treap.Insert(4)
	treap.Insert(5)
	fmt.Println("aisfosijf")
	fmt.Println(treap)
	fmt.Println(treap.root.value, treap.root.left.value, treap.root.right.value)
}
