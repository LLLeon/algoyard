package queue

/*
 ArrayQueue
*/

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

/*
 CircularQueue
*/

// CircularQueue represent a circular queue.
// When the queue is full, no data is stored at the tail location,
// so the circular queue wastes one storage space
type CircularQueue struct {
	items []int
	size  int
	head  int
	tail  int
}

// NewCircularQueue returns an initialized CircularQueue.
func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		items: make([]int, size),
		size:  size,
		head:  0,
		tail:  0,
	}
}

// Enqueue the item.
func (cq *CircularQueue) Enqueue(item int) bool {
	// queue is full
	if (cq.tail+1)%cq.size == cq.head {
		return false
	}

	cq.items[cq.tail] = item
	cq.tail = (cq.tail + 1) % cq.size

	return true
}

// Dequeue returns an item from the queue.
func (cq *CircularQueue) Dequeue() (int, bool) {
	// queue is empty
	if cq.head == cq.tail {
		return 0, false
	}

	item := cq.items[cq.head]
	cq.head = (cq.head + 1) % cq.size

	return item, true
}
