package treap

import (
	"testing"
)

func TestTreap(t *testing.T) {
	treap := NewTreap()
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
	treap := NewTreap()
	treap.Insert(1)
	treap.Insert(2)
	treap.Insert(3)
	treap.Insert(4)
	treap.Insert(5)
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
