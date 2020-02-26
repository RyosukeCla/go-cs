package linkedlist

import (
	"testing"
)

func TestInsert(t *testing.T) {
	list := New()
	one := list.InsertBiginning(1)
	two := list.InsertBiginning(2)
	three := list.InsertEnd(3)

	// 2 <-> 1 <-> 3

	if list.Len() != 3 {
		t.Fatal("error")
	}
	if three.Next() != nil || three.Prev() != one {
		t.Fatal("error")
	}
	if one.Next() != three || one.Prev() != two {
		t.Fatal("error")
	}
}

func TestSwap(t *testing.T) {
	list := New()
	one := list.InsertBiginning(1)
	two := list.InsertBiginning(2)
	three := list.InsertEnd(3)
	four := list.InsertEnd(4)

	// 2 <-> 1 <-> 3 <-> 4

	list.Swap(one, four)

	// 2 <-> 4 <-> 3 <-> 1
	if two.Next() != four || three.Next() != one {
		t.Fatal("error")
	}
	if one.Prev() != three || one.Next() != nil || list.Last() != one {
		t.Fatal("error")
	}
	if four.Next() != three {
		t.Fatal("error")
	}
}

func TestMoveToBeginning(t *testing.T) {
	list := New()
	one := list.InsertBiginning(1)
	two := list.InsertBiginning(2)
	three := list.InsertEnd(3)

	// 2 <-> 1 <-> 3
	list.MoveToBeginning(three)
	// 3 <-> 2 <-> 1

	if three.Next() != two || three.Prev() != nil || list.First() != three {
		t.Fatal("error")
	}
	if one.Next() != nil || list.Last() != one {
		t.Fatal("error")
	}
}

func TestMoveToEnd(t *testing.T) {
	list := New()
	one := list.InsertBiginning(1)
	two := list.InsertBiginning(2)
	three := list.InsertEnd(3)

	// 2 <-> 1 <-> 3
	list.MoveToEnd(two)
	// 1 <-> 3 <-> 2

	if three.Next() != two || two.Prev() != three || two.Next() != nil || list.Last() != two {
		t.Fatal("error")
	}
	if one.Prev() != nil || list.First() != one {
		t.Fatal("error")
	}
}

func TestRemove(t *testing.T) {
	list := New()
	one := list.InsertBiginning(1)
	two := list.InsertBiginning(2)
	three := list.InsertEnd(3)

	// 2 <-> 1 <-> 3
	list.Remove(one)
	// 2 <-> 3

	if list.Len() != 2 {
		t.Fatal("error")
	}

	if two.Next() != three || three.Prev() != two {
		t.Fatal("error")
	}
}
