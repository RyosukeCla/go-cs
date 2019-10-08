package adt

import (
	"fmt"
	"testing"
)

// [29 19 9 4 10 1 2 3 2 2]
func TestMaxHeap(t *testing.T) {
	arr := []float32{4, 2, 9, 19, 29, 1, 2, 3, 2, 10}

	heap := NewHeap("max", 20)

	for _, element := range arr {
		heap.Insert(element)
	}

	fmt.Println(heap)
}

func TestMinHeap(t *testing.T) {
	arr := []float32{4, 2, 9, 19, 29, 1, 2, 3, 2, 10}

	heap := NewHeap("min", 20)

	for _, element := range arr {
		heap.Insert(element)
	}

	fmt.Println(heap)
}
