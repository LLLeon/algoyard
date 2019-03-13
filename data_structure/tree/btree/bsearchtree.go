package btree

// Find implements a binary search tree search for specified data.
func Find(root *Node, data int) *Node {
	if root.data == data {
		return root
	}

	p := root
	for p != nil {
		if data < p.data {
			p = p.left
		} else if data > p.data {
			p = p.right
		} else {
			return p
		}
	}

	return nil
}

// Insert the specified data into the binary search tree.
func Insert(root *Node, data int) *Node {
	if root == nil {
		root = NewNode(data)
		return root
	}

	p := root
	for p != nil {
		if data > p.data {
			if p.right == nil {
				p.right = NewNode(data)
				return root
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = NewNode(data)
				return root
			}
			p = p.left
		}
	}

	return root
}

// Delete implements the removal of specified data from
// the binary search tree.
func Delete(root *Node, data int) {
	if root == nil {
		return
	}

	var pp *Node // the parent node of p
	p := root    // the Node to delete

	// find the node to delete
	for p != nil && p.data != data {
		pp = p
		if data > p.data {
			p = p.right
		} else {
			p = p.left
		}
	}

	// not found
	if p == nil {
		return
	}

	// the node to be deleted has two child nodes.
	// find the smallest node in the right subtree.
	if p.left != nil && p.right != nil {
		minP := p.right
		minPP := p

		for minP.left != nil {
			minPP = minP
			minP = minP.left
		}

		// Replace the value of p with the value of minP.
		// The node to delete becomes minP.
		p.data = minP.data
		p = minP
		pp = minPP
	}

	// The node to be deleted is a leaf node or has only
	// one child node.
	var child *Node

	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	// The node to delete is the root node.
	if pp == nil {
		root = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}
}

// MaxNode returns the largest node in the binary search tree.
func MaxNode(root *Node) *Node {
	if root == nil {
		return nil
	}

	max := root

	for max.right != nil {
		max = max.right
	}

	return max
}

// MinNode returns the minimum node in the binary search tree.
func MinNode(root *Node) *Node {
	if root == nil {
		return nil
	}

	min := root

	for min.left != nil {
		min = min.left
	}

	return min
}
