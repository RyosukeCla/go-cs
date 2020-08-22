package queue

import (
	"testing"
)

func Test(t *testing.T) {
	list := New()
	list.Enqueue(0)
	list.Enqueue(1)
	list.Enqueue(2)
	if list.Len() != 3 {
		t.Fatal("error")
	}

	value1 := list.Dequeue()
	if value1 != 0 {
		t.Fatal("error")
	}
	list.Dequeue()
	list.Dequeue()
	if list.Len() != 0 {
		t.Fatal("error")
	}
}
