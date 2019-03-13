package binary

import "fmt"

// Node represents a node in a binary tree.
type Node struct {
	data  int
	left  *Node
	right *Node
}

// NewNode returns a new Node with data.
func NewNode(data int) *Node {
	return &Node{data: data}
}

func (n *Node) String() string {
	return fmt.Sprintf("value: %+v, left: %+v, right: %+v\n",
		n.data, n.left, n.right)
}
