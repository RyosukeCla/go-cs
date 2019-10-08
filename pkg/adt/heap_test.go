package adt

import (
	"fmt"
	"testing"
)

// [29 19 9 4 10 1 2 3 2 2]
func TestHeap(t *testing.T) {
	arr := []float32{4, 2, 9, 19, 29, 1, 2, 3, 2, 10}
	n := len(arr)

	start := (n - 1) / 2
	for i := start; i >= 0; i-- {
		maxHeapify(arr, i, n)
	}

	fmt.Println(arr)
}

// [29 19 9 4 10 1 2 3 2 2]
func TestMaxHeap(t *testing.T) {
	arr := []float32{4, 2, 9, 19, 29, 1, 2, 3, 2, 10}

	heap := NewMaxHeap(20)

	for _, element := range arr {
		heap.Insert(element)
	}

	fmt.Println(heap)
}

func TestMinHeap(t *testing.T) {
	arr := []float32{4, 2, 9, 19, 29, 1, 2, 3, 2, 10}

	heap := NewMinHeap(20)

	for _, element := range arr {
		heap.Insert(element)
	}

	fmt.Println(heap)
}
