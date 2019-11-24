package binary

// Search searchs index of target value in sorted array.
func Search(array []int, target int) int {
	left := 0
	right := len(array) - 1

	for right-left >= 0 {
		middle := (right + left) / 2
		element := array[middle]

		if element == target {
			return middle
		} else if element < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
