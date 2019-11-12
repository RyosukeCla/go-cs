package linear

// Search returns index
func Search(array []int, target int) int {
	for i, element := range array {
		if target == element {
			return i
		}
	}

	return -1
}
