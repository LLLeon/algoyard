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

func TestFindKthTail(t *testing.T) {
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

	k := 1

	kthNeg1, existNeg1 := FindKthTail(head, -1)
	if existNeg1 {
		fmt.Printf("the reciprocal 1th node: %d\n", kthNeg1.Value)
	} else {
		fmt.Printf("the reciprocal %dth node does not exist.\n", -1)
	}

	kth0, exist1 := FindKthTail(head, 0)
	if exist1 {
		fmt.Printf("the reciprocal 1th node: %d\n", kth0.Value)
	} else {
		fmt.Printf("the reciprocal %dth node does not exist.\n", 0)
	}

	kth1, exist1 := FindKthTail(head, k)
	if exist1 {
		k++
		fmt.Printf("the reciprocal 1th node: %d\n", kth1.Value)
	} else {
		fmt.Printf("the reciprocal %dth node does not exist.\n", k)
	}

	kth2, exist2 := FindKthTail(head, k)
	if exist2 {
		k++
		fmt.Printf("the reciprocal 2th node: %d\n", kth2.Value)
	} else {
		fmt.Printf("the reciprocal %dth node does not exist.\n", k)
	}

	kth3, exist3 := FindKthTail(head, k)
	if exist3 {
		k++
		fmt.Printf("the reciprocal 3th node: %d\n", kth3.Value)
	} else {
		fmt.Printf("the reciprocal %dth node does not exist.\n", k)
	}

	kth4, exist4 := FindKthTail(head, k)
	if exist4 {
		k++
		fmt.Printf("the reciprocal 4th node: %d\n", kth4.Value)
	} else {
		fmt.Printf("the reciprocal %dth node does not exist.\n", k)
	}

	kth5, exist5 := FindKthTail(head, k)
	if exist5 {
		k++
		fmt.Printf("the reciprocal 4th node: %d\n", kth5.Value)
	} else {
		fmt.Printf("the reciprocal %dth node does not exist.\n", k)
	}
}
