package jump

import (
	"math"
)

// Search searches an element from sorted array
func Search(array []int, target int) int {
	n := len(array)
	blockSize := int(math.Sqrt(float64(n)))
	index := 0
	for index < n {
		if array[index] >= target {
			for back := 0; back < blockSize; back++ {
				if array[index-back] == target {
					return index - back
				}
			}
			return -1
		}
		index += blockSize
	}

	rest := n - index + blockSize
	for back := 1; back <= rest; back++ {
		if array[n-back] == target {
			return n - back
		}
	}

	return -1
}
