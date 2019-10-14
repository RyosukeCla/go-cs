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
	fmt.Println(treap)
	fmt.Println(treap.root.key, treap.root.left.key, treap.root.right.key)
}

func TestTreap_Remove(t *testing.T) {
	treap := NewTreap()
	treap.Insert(1)
	treap.Insert(2)
	treap.Insert(3)
	treap.Insert(4)
	treap.Insert(5)
	treap.Remove(3)
	treap.Remove(2)

	fmt.Println("addf-1")
	fmt.Println("asdf-2", treap.Search(2), treap.Search(4))
	fmt.Println("asdf-3", treap.Search(3), treap.Search(5))
	fmt.Println("asdf-4", treap.Search(3), treap.Search(100))
}
