package queue

// ArrayQueue represents a queue implemented with an array.
type ArrayQueue struct {
	items []int
	size  int
	head  int
	tail  int
}

// NewArrayQueue returns an initialized ArrayQueue.
func NewArrayQueue(size int) *ArrayQueue {
	return &ArrayQueue{
		items: make([]int, size),
		size:  size,
		head:  0,
		tail:  0,
	}
}

// Enqueue the item.
func (aq *ArrayQueue) Enqueue(item int) bool {
	if aq.tail == aq.size {
		if aq.head == 0 {
			return false
		}

		for i := aq.head; i < aq.tail; i++ {
			aq.items[i-aq.head] = aq.items[i]
		}

		aq.tail = aq.tail - aq.head
		aq.head = 0
	}

	aq.items[aq.tail] = item
	aq.tail++

	return true
}

// Dequeue returns an item from the queue.
func (aq *ArrayQueue) Dequeue() (int, bool) {
	if aq.head == aq.tail {
		return 0, false
	}

	item := aq.items[aq.head]
	aq.head++

	return item, true
}
