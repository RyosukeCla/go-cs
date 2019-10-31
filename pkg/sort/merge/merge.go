package merge

/*
1. Find the middle point to divide the array into two halves:
	middle m = (l+r)/2
2. Call mergeSort for first half:
	Call mergeSort(arr, l, m)
3. Call mergeSort for second half:
	Call mergeSort(arr, m1, r)
4. Merge the two halves sorted in step 2 and 3:
	Call merge(arr, l, m, r)

	from https://www.geeksforgeeks.org/merge-sort/
*/

func merge(array []int, left int, middle int, right int) {
	leftN := middle - left
	rightN := right - middle
	leftArray := make([]int, leftN)
	rightArray := make([]int, rightN)
	for i := 0; i < leftN; i++ {
		leftArray[i] = array[left+i]
	}
	for i := 0; i < rightN; i++ {
		rightArray[i] = array[middle+i]
	}

	leftCursor := 0
	rightCursor := 0

	for i := left; i < right; i++ {
		if leftCursor >= leftN {
			array[i] = rightArray[rightCursor]
			rightCursor++
			continue
		}
		if rightCursor >= rightN {
			array[i] = leftArray[leftCursor]
			leftCursor++
			continue
		}

		if leftArray[leftCursor] < rightArray[rightCursor] {
			array[i] = leftArray[leftCursor]
			leftCursor++
		} else {
			array[i] = rightArray[rightCursor]
			rightCursor++
		}
	}
}

func mergeSort(array []int, left int, right int) {
	if right-left > 1 {
		middle := (left + right) / 2
		mergeSort(array, left, middle)
		mergeSort(array, middle, right)
		merge(array, left, middle, right)
	}
}

// Sort sorts an array
func Sort(array []int) {
	mergeSort(array, 0, len(array))
}
