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
		head:  -1,
		tail:  -1,
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
		head:  -1,
		tail:  -1,
	}
}

// Enqueue the item.
func (cq *CircularQueue) Enqueue(item int) bool {
	// queue is full
	if cq.IsFull() {
		return false
	}

	if cq.IsEmpty() {
		cq.head = 0
	}

	cq.tail = (cq.tail + 1) % cq.size
	cq.items[cq.tail] = item

	return true
}

// Dequeue delete an item from the queue.
func (cq *CircularQueue) Dequeue() bool {
	// queue is empty
	if cq.IsEmpty() {
		return false
	}

	// the queue is empty
	if cq.head == cq.tail {
		cq.head = -1
		cq.tail = -1

		return true
	}

	cq.head = (cq.head + 1) % cq.size

	return true
}

// Front get the front item from the queue.
func (cq *CircularQueue) Front() int {
	if cq.IsEmpty() {
		return -1
	}

	return cq.items[cq.head]
}

// Rear get the last item from the queue.
func (cq *CircularQueue) Rear() int {
	if cq.IsEmpty() {
		return -1
	}

	return cq.items[cq.tail]
}

func (cq *CircularQueue) IsEmpty() bool {
	return cq.head == -1
}

func (cq *CircularQueue) IsFull() bool {
	return (cq.tail+1)%cq.size == cq.head
}
