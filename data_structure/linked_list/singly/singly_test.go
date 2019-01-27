package singly

import (
	"fmt"
	"testing"
)

func TestConstructList(t *testing.T) {
	values := []int{1, 2, 3}

	head := ConstructList(values)
	if head.Value == 0 {
		fmt.Printf("head: %d\n", head.Value)
	} else {
		fmt.Printf("[error] head want: %d, get: %d\n", 0, head.Value)
	}

	node1 := head.Next
	if node1.Value == 3 {
		fmt.Printf("node1: %d\n", node1.Value)
	} else {
		fmt.Printf("[error] node1 want: %d, get: %d\n", 3, node1.Value)
	}

	node2 := node1.Next
	if node2.Value == 2 {
		fmt.Printf("node2: %d\n", node2.Value)
	} else {
		fmt.Printf("[error] node2 want: %d, get: %d\n", 2, node2.Value)
	}

	node3 := node2.Next
	if node3.Value == 1 {
		fmt.Printf("node3: %d\n", node3.Value)
	} else {
		fmt.Printf("[error] node3 want: %d, get: %d\n", 1, node2.Value)
	}
}

func TestConstructListPositiveSeq(t *testing.T) {
	values := []int{1, 2, 3}

	head := ConstructListPositiveSeq(values)
	if head.Value == 0 {
		fmt.Printf("head: %d\n", head.Value)
	} else {
		fmt.Printf("[error] head want: %d, get: %d\n", 0, head.Value)
	}

	node1 := head.Next
	if node1.Value == 1 {
		fmt.Printf("node1: %d\n", node1.Value)
	} else {
		fmt.Printf("[error] node1 want: %d, get: %d\n", 1, node1.Value)
	}

	node2 := node1.Next
	if node2.Value == 2 {
		fmt.Printf("node2: %d\n", node2.Value)
	} else {
		fmt.Printf("[error] node2 want: %d, get: %d\n", 2, node2.Value)
	}

	node3 := node2.Next
	if node3.Value == 3 {
		fmt.Printf("node3: %d\n", node3.Value)
	} else {
		fmt.Printf("[error] node3 want: %d, get: %d\n", 3, node2.Value)
	}
}

func TestReverse(t *testing.T) {
	values := []int{1, 2, 3}

	// 3 -> 2 -> 1
	head := ConstructList(values)

	fmt.Printf("head: %d\n", head.Value)

	node1 := head.Next
	fmt.Printf("node1: %d\n", node1.Value)

	node2 := node1.Next
	fmt.Printf("node2: %d\n", node2.Value)

	node3 := node2.Next
	fmt.Printf("node3: %d\n", node3.Value)

	// 1 -> 2 -> 3 -> 0
	newHead := Reverse(head)
	node1 = newHead
	if node1.Value == 1 {
		fmt.Printf("new node1: %d\n", node1.Value)
	} else {
		fmt.Printf("[error] new node1 want: %d, get: %d\n", 3, node1.Value)
	}

	node2 = node1.Next
	if node2.Value == 2 {
		fmt.Printf("new node2: %d\n", node2.Value)
	} else {
		fmt.Printf("[error] new node2 want: %d, get: %d\n", 2, node2.Value)
	}

	node3 = node2.Next
	if node3.Value == 3 {
		fmt.Printf("new node3: %d\n", node3.Value)
	} else {
		fmt.Printf("[error] new node3 want: %d, get: %d\n", 3, node2.Value)
	}

	fmt.Printf("the original head node: %d\n", node3.Next.Value)
}
