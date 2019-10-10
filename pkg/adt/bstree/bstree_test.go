package bstree

import (
	"fmt"
	"testing"
)

func TestBSTree(t *testing.T) {
	tree := NewBSTree()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(30)

	fmt.Println(tree.Search(20))
	fmt.Println(tree.Search(100))

	tree.Remove(20)
	fmt.Println(tree.Search(30))
}
