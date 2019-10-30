package heap

type compare func(a int, b int) bool

// Heap is a heap
type Heap struct {
	array   []int
	compare compare
}

// NewHeap returns heap. heapType is "max" or "min"
func NewHeap(heapType string, cap int) Heap {
	array := make([]int, 0, cap)

	var compare compare
	if heapType == "min" {
		compare = func(a int, b int) bool {
			return a < b
		}
	} else {
		compare = func(a int, b int) bool {
			return a > b
		}
	}
	return Heap{
		array:   array,
		compare: compare,
	}
}

// Insert inserts element with keeping max heap property
func (h *Heap) Insert(element int) {
	h.array = append(h.array, element)
	current := len(h.array) - 1
	for {
		parent := parentIndex(current)
		if h.compare(h.array[current], h.array[parent]) {
			h.array[current], h.array[parent] = h.array[parent], h.array[current]
			current = parent
		} else {
			break
		}
	}
}

// Extract extracts max/min element with keeping max/min heap property
func (h *Heap) Extract() int {
	first := h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1] // remove first
	heapify(h.array, h.compare, 0, len(h.array))
	return first
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftChildIndex(i int) int {
	return 2*i + 1
}

func rightChildIndex(i int) int {
	return 2*i + 2
}

func heapify(arr []int, compare compare, i int, size int) {
	left := leftChildIndex(i)
	right := rightChildIndex(i)
	largest := i

	// If left child is larger than root
	if left < size && compare(arr[left], arr[largest]) {
		largest = left
	}

	// If right child is larger than largest so far
	if right < size && compare(arr[right], arr[largest]) {
		largest = right
	}

	// If largest is not root
	if largest != i {
		// swap
		arr[i], arr[largest] = arr[largest], arr[i]

		// Recursively heapify the affected sub-tree
		heapify(arr, compare, largest, size)
	}
}
