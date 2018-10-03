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
func Remove(head, toDelete *ListNode) {
	if toDelete == nil {
		return
	}

	// if toDelete is an intermediate node
	if toDelete.Next != nil {
		oldNext := toDelete.Next

		toDelete.Value = oldNext.Value
		toDelete.Next = oldNext.Next

		oldNext = nil
	} else {
		// if only one node
		if head == toDelete {
			head = nil
		} else {
			node := head

			// find the next-to-last node
			for {
				if node.Next == toDelete {
					break
				}
				node = node.Next
			}

			node.Next = nil
		}

		toDelete = nil
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
