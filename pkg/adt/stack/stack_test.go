package stack

import (
	"testing"
)

func Test(t *testing.T) {
	list := New()
	list.Push(0)
	list.Push(1)
	list.Push(2)
	if list.Len() != 3 {
		t.Fatal("error")
	}

	value1 := list.Pop()
	if value1 != 2 {
		t.Fatal("error")
	}
	list.Pop()
	list.Pop()
	if list.Len() != 0 {
		t.Fatal("error")
	}
}
