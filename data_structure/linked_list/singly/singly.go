package singly

// ListNode represents a node of a singly linked list.
type ListNode struct {
	Value int
	Next  *ListNode
}

// NewListNode creates a new node with val.
func NewListNode(val int) *ListNode {
	return &ListNode{
		Value: val,
	}
}

// NextNode returns the node after the ln.
func (ln *ListNode) NextNode() *ListNode {
	return ln.Next
}

// InsertAfter inserts the node with val after ln and return val.
func (ln *ListNode) InsertAfter(val int) int {
	return insert(val, ln)
}

// InsertAt inserts the node with val after at and return the node.
func InsertAt(val int, at *ListNode) *ListNode {
	insert(val, at)
	return at.Next
}

// Remove removes the specified node.
func Remove(node *ListNode) {
	if node == nil {
		return
	}

	if node.Next != nil {
		oldNext := node.Next

		node.Value = oldNext.Value
		node.Next = oldNext.Next

		oldNext.Value = 0
		oldNext.Next = nil
	} else {
		node.Value = 0
		node.Next = nil
	}
}

func insert(val int, at *ListNode) int {
	if at.Next == nil {
		next := &ListNode{}
		next.Value = val
		at.Next = next
	} else {
		oldNext := at.Next
		newNext := &ListNode{}

		newNext.Value = val
		newNext.Next = oldNext

		at.Next = newNext
	}

	return val
}
