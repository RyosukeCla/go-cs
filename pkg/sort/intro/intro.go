package intro

import (
	"math"

	"github.com/RyosukeCla/go-cs/pkg/sort/heap"
	"github.com/RyosukeCla/go-cs/pkg/sort/insertion"
	"github.com/RyosukeCla/go-cs/pkg/sort/quick"
)

func introSort(array []int, depthLimit int) {
	n := len(array)
	if n <= 16 {
		insertion.Sort(array)
	} else if depthLimit == 0 {
		heap.Sort(array)
	} else {
		p := quick.Partition(array, 0, n-1)
		introSort(array[0:p], depthLimit-1)
		introSort(array[p+1:n], depthLimit-1)
	}
}

// Sort sorts an array with introspective sort algorithm
func Sort(array []int) {
	n := float64(len(array))
	depthLimit := 2 * int(math.Floor(math.Log2(n)))
	introSort(array, depthLimit)
}
