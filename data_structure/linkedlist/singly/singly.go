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
		Next:  nil,
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

// ConstructList builds the linked list using a head interpolation method.
func ConstructList(values []int) *ListNode {
	head := &ListNode{Value: 0, Next: nil}

	for i := 0; i < len(values); i++ {
		node := &ListNode{Value: values[i], Next: nil}
		node.Next = head.Next
		head.Next = node
	}

	return head
}

// constructListPositiveSeq builds the linked list using a head interpolation
// method but with the positive sequence.
func ConstructListPositiveSeq(values []int) *ListNode {
	head := &ListNode{Value: 0, Next: nil}

	for i := len(values) - 1; i >= 0; i-- {
		node := &ListNode{Value: values[i], Next: nil}
		node.Next = head.Next
		head.Next = node
	}

	return head
}

// Reverse a linked list and return the head node of the linked list.
func Reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *ListNode = nil
	curr := head

	for {
		if curr == nil {
			break
		}

		temp := curr     // 1: record current node
		curr = curr.Next // 2: move the current node backwards
		temp.Next = prev // 3: point Next of the old current node to the previous node
		prev = temp      // 4: assign the prev to the old current node
	}

	// returns the last node as the head of the reversed linked list.
	return prev
}

func ReverseRecursive(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	last := ReverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}

// FindKthTail returns the reciprocal Kth node.
func FindKthTail(head *ListNode, k int) (*ListNode, bool) {
	if head == nil || k <= 0 {
		return nil, false
	}

	fastP := head
	slowP := head

	for i := 0; i < k; i++ {
		if fastP != nil {
			fastP = fastP.Next
		} else {
			return nil, false
		}
	}

	for fastP != nil {
		fastP = fastP.Next
		slowP = slowP.Next
	}

	return slowP, true
}

// GetRingLength returns the length of the ring in a linked list.
func GetRingLength(head *ListNode) int {
	slowP := head
	fastP := head

	for slowP != nil && fastP.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next

		// the two pointers meet for the first time
		if slowP == fastP {
			break
		}
	}

	if slowP == nil || fastP.Next == nil {
		return -1
	}

	slowP = slowP.Next
	fastP = fastP.Next.Next
	length := 1

	// loop to meet again
	for slowP != fastP {
		slowP = slowP.Next
		fastP = fastP.Next.Next
		length++
	}

	return length
}
