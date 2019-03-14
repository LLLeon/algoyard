package heap

// Heap represents the heap implemented in an array.
// The first position of the array does not store element.
type Heap struct {
	items []int
	count int
	size  int
}

func NewHeap(size int) *Heap {
	return &Heap{
		items: make([]int, size+1),
		count: 0,
		size:  size,
	}
}

// Insert implements inserting elements into the heap.
func (h *Heap) Insert(item int) bool {
	if h.count >= h.size {
		return false
	}

	h.count++
	h.items[h.count] = item
	i := h.count

	// Adjust the location of nodes and their children.
	// Use large elements as parent nodes.
	for (i/2 > 0) && (h.items[i] > h.items[i/2]) {
		h.items[i], h.items[i/2] = h.items[i/2], h.items[i]
		i = i / 2
	}

	return true
}

// RemoveMax implements remove an element from the heap.
func (h *Heap) RemoveMax() bool {
	if h.count == 0 {
		return false
	}

	h.items[1] = h.count
	h.count--

	heapify(h.items, h.count, 1)
	return true
}

// heapify from the top down.
func heapify(arr []int, count, i int) {
	for {
		maxPos := i

		if (i*2 <= count) && (arr[i] < arr[i*2]) {
			maxPos = i * 2
		}
		if (i*2+1 <= count) && (arr[maxPos] < arr[i*2+1]) {
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}

		arr[i], arr[maxPos] = arr[maxPos], arr[i]
		i = maxPos
	}
}
