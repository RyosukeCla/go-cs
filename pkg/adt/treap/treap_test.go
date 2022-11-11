package treap

import (
	"testing"

	"github.com/RyosukeCla/go-cs/pkg/adt/bstree"
)

func TestTreap(t *testing.T) {
	treap := NewTreap(bstree.NumberComparator[int])
	treap.Insert(1)
	treap.Insert(2)
	treap.Insert(3)
	treap.Insert(4)
	treap.Insert(5)

	if treap.Search(5) == false {
		t.Fatal("error")
	}
}

func TestTreap_Remove(t *testing.T) {
	treap := NewTreap(bstree.NumberComparator[int])
	treap.Insert(1)
	treap.Insert(2)
	treap.Insert(3)
	treap.Insert(4)
	treap.Insert(5)

	if treap.Search(3) == false {
		t.Fatal("error")
	}

	if treap.Search(2) == false {
		t.Fatal("error")
	}

	treap.Remove(3)
	treap.Remove(2)

	if treap.Search(3) == true {
		t.Fatal("error")
	}

	if treap.Search(2) == true {
		t.Fatal("error")
	}

	if treap.Search(5) == false {
		t.Fatal("error")
	}
}
