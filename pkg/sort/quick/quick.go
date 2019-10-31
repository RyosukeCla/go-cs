package quick

// returns partition index and partition into >= pivot and < pivot.
func partition(array []int, left int, right int) int {
	pivot := array[left] // select first element as pivot
	i := left + 1

	// For every element in array[left+1:right] excluding pivot with left scanning.
	for j := left + 1; j <= right; j++ {
		// If j'th element is lower than pivot.
		if array[j] < pivot {
			// Swap j'th element with i'th element.
			array[i], array[j] = array[j], array[i]
			i++
		}
	}

	array[left], array[i-1] = array[i-1], array[left]
	return i - 1
}

func quickSort(array []int, left int, right int) {
	if left < right {
		partitionIndex := partition(array, left, right)
		quickSort(array, left, partitionIndex-1)
		quickSort(array, partitionIndex+1, right)
	}
}

// Sort sorts an array with quick sorting algorithm
func Sort(array []int) {
	quickSort(array, 0, len(array)-1)
}
