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

// Push an item to the top of the stack, when the maximum
// capacity is reached, capacity to double.
func (as *ArrayStack) Push(item string) bool {
	if as.count == as.size {
		newItems := make([]string, as.size*2)

		for i, item := range as.items {
			newItems[i] = item
		}

		as.items = newItems
		as.size = as.size * 2
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

// Top returns an item from the top of the stack.
func (as *ArrayStack) Top() (string, bool) {
	if as.count == 0 {
		return "", false
	}

	item := as.items[as.count-1]
	return item, true
}

// Size returns the size of the stace.
func (as *ArrayStack) Size() int64 {
	return as.size
}
