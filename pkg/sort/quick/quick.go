package quick

/*

Partitioning Algorithm

Supporse,
  - array[left] = pivot
  - array[left+1:i] < pivot
  - array[i:j] >= pivot
  - where j >= i >= left+1

If array[j] < pivot and swap array[j] and array[i],
Then,
  - array[left+1:i+1] < pivot
  - array[i+1:j+1] >= pivot
And update i <- i+1, j <- j+1 remaining these properties
Then do until j=right

Lastly, by swapping array[left] and array[i+1], we get
  - array[left:i] < pivot
  - array[i:right] >= pivot

*/

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

/*
Quick Sort Algorithm

Suppose,
  - array A

By adopting partitioning, we get two subarrays A1, A2 which fulfills
  - a_1i <= a_2j for all i, j (since a_1i <= pivot <= a_2j)
In the same way, we split subarrays into two arrays until end.
Then we get a_ik <= a_jl for all i < j, k < l.
Thus a_i <= a_j for all i < j where a_i,j in A,
which is the result of sort.

By adopting partitioning, we get a_ik < a_il and pivot_i < pivot_j
where for all (k, l)'th elements in (i, j)'th arrays of subarrays.
Since a_i0 < pivot_i < a_iN < pivot_i+1 where N is length of a_i array,
we get a_ik < a_jl for all i, j, k, l.

*/
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
