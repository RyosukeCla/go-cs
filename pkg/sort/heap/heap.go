package heap

import (
	"github.com/RyosukeCla/go-cs/pkg/adt/heap"
)

// Sort sorts an array with heap sorting algorithm
func Sort(array []int) {
	n := len(array)
	heap := heap.NewHeap("min", len(array))

	for _, element := range array {
		heap.Insert(element)
	}

	for i := 0; i < n; i++ {
		array[i] = heap.Extract()
	}
}
