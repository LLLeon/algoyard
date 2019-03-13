package binary

import "fmt"

// PreOrder implements pre-order traversal of binary tree.
// Each node is accessed up to twice,
// time complexity is O(n).
func PreOrder(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.data)
	PreOrder(root.left)
	PreOrder(root.right)
}

// InOrder implements in-order traversal of binary tree.
// Each node is accessed up to twice,
// time complexity is O(n).
func InOrder(root *Node) {
	if root == nil {
		return
	}

	InOrder(root.left)
	fmt.Println(root.data)
	InOrder(root.right)
}

// PostOrder implements post-order traversal of binary tree.
// Each node is accessed up to twice,
// time complexity is O(n).
func PostOrder(root *Node) {
	if root == nil {
		return
	}

	PostOrder(root.left)
	PostOrder(root.right)
	fmt.Println(root.data)
}
