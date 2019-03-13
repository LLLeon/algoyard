package btree

// BinaryTree represents a binary tree.
type BinaryTree struct {
	root *Node
}

// NewBinaryTree returns a new BinaryTree with root value.
func NewBinaryTree(rootVal int) *BinaryTree {
	return &BinaryTree{NewNode(rootVal)}
}
