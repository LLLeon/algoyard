package stack

// ArrayStack represents a stack implemented with array.
type ArrayStack struct {
	items []string
	count int64
	size  int64
}

// NewArrayStack returns a new ArrayStack with specified size.
func NewArrayStack(size int64) *ArrayStack {
	return &ArrayStack{
		items: make([]string, size),
		count: 0,
		size:  size,
	}
}

// Push an item to the top of the stack.
func (as *ArrayStack) Push(item string) bool {
	if as.count == as.size {
		return false
	}

	as.items[as.count] = item
	as.count++

	return true
}

// Pop pops up an item from the top of the stack.
func (as *ArrayStack) Pop() (string, bool) {
	if as.count == 0 {
		return "", false
	}

	item := as.items[as.count-1]
	as.count--

	return item, true
}
