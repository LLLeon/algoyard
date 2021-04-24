package singly

import (
	"testing"
)

func TestConstructList(t *testing.T) {
	values := []int{1, 2, 3}

	head := ConstructList(values)
	if head.Value == 0 {
		t.Logf("head: %d\n", head.Value)
	} else {
		t.Errorf("[error] head want: %d, get: %d\n", 0, head.Value)
	}

	node1 := head.Next
	if node1.Value == 3 {
		t.Logf("node1: %d\n", node1.Value)
	} else {
		t.Errorf("[error] node1 want: %d, get: %d\n", 3, node1.Value)
	}

	node2 := node1.Next
	if node2.Value == 2 {
		t.Logf("node2: %d\n", node2.Value)
	} else {
		t.Errorf("[error] node2 want: %d, get: %d\n", 2, node2.Value)
	}

	node3 := node2.Next
	if node3.Value == 1 {
		t.Logf("node3: %d\n", node3.Value)
	} else {
		t.Errorf("[error] node3 want: %d, get: %d\n", 1, node2.Value)
	}
}

func TestConstructListPositiveSeq(t *testing.T) {
	values := []int{1, 2, 3}

	head := ConstructListPositiveSeq(values)
	if head.Value == 0 {
		t.Logf("head: %d\n", head.Value)
	} else {
		t.Errorf("[error] head want: %d, get: %d\n", 0, head.Value)
	}

	node1 := head.Next
	if node1.Value == 1 {
		t.Logf("node1: %d\n", node1.Value)
	} else {
		t.Errorf("[error] node1 want: %d, get: %d\n", 1, node1.Value)
	}

	node2 := node1.Next
	if node2.Value == 2 {
		t.Logf("node2: %d\n", node2.Value)
	} else {
		t.Errorf("[error] node2 want: %d, get: %d\n", 2, node2.Value)
	}

	node3 := node2.Next
	if node3.Value == 3 {
		t.Logf("node3: %d\n", node3.Value)
	} else {
		t.Errorf("[error] node3 want: %d, get: %d\n", 3, node2.Value)
	}
}

func TestReverse(t *testing.T) {
	values := []int{1, 2, 3}

	// 3 -> 2 -> 1
	head := ConstructList(values)

	t.Logf("head: %d\n", head.Value)

	node1 := head.Next
	t.Logf("node1: %d\n", node1.Value)

	node2 := node1.Next
	t.Logf("node2: %d\n", node2.Value)

	node3 := node2.Next
	t.Logf("node3: %d\n", node3.Value)

	// 1 -> 2 -> 3 -> 0
	newHead := Reverse(head)
	node1 = newHead
	if node1.Value == 1 {
		t.Logf("new node1: %d\n", node1.Value)
	} else {
		t.Errorf("[error] new node1 want: %d, get: %d\n", 3, node1.Value)
	}

	node2 = node1.Next
	if node2.Value == 2 {
		t.Logf("new node2: %d\n", node2.Value)
	} else {
		t.Errorf("[error] new node2 want: %d, get: %d\n", 2, node2.Value)
	}

	node3 = node2.Next
	if node3.Value == 3 {
		t.Logf("new node3: %d\n", node3.Value)
	} else {
		t.Errorf("[error] new node3 want: %d, get: %d\n", 3, node2.Value)
	}

	t.Logf("the original head node: %d\n", node3.Next.Value)
}

func TestReverseRecursive(t *testing.T) {
	values := []int{1, 2, 3}

	// 3 -> 2 -> 1
	head := ConstructList(values)

	t.Logf("head: %d\n", head.Value)

	node1 := head.Next
	t.Logf("node1: %d\n", node1.Value)

	node2 := node1.Next
	t.Logf("node2: %d\n", node2.Value)

	node3 := node2.Next
	t.Logf("node3: %d\n", node3.Value)

	// 1 -> 2 -> 3 -> 0
	newHead := ReverseRecursive(head)
	node1 = newHead
	if node1.Value == 1 {
		t.Logf("new node1: %d\n", node1.Value)
	} else {
		t.Errorf("[error] new node1 want: %d, get: %d\n", 3, node1.Value)
	}

	node2 = node1.Next
	if node2.Value == 2 {
		t.Logf("new node2: %d\n", node2.Value)
	} else {
		t.Errorf("[error] new node2 want: %d, get: %d\n", 2, node2.Value)
	}

	node3 = node2.Next
	if node3.Value == 3 {
		t.Logf("new node3: %d\n", node3.Value)
	} else {
		t.Errorf("[error] new node3 want: %d, get: %d\n", 3, node2.Value)
	}

	t.Logf("the original head node: %d\n", node3.Next.Value)
}

func TestFindKthTail(t *testing.T) {
	values := []int{1, 2, 3}

	// 3 -> 2 -> 1
	head := ConstructList(values)

	t.Logf("head: %d\n", head.Value)

	node1 := head.Next
	t.Logf("node1: %d\n", node1.Value)

	node2 := node1.Next
	t.Logf("node2: %d\n", node2.Value)

	node3 := node2.Next
	t.Logf("node3: %d\n", node3.Value)

	k := 1

	kthNeg1, existNeg1 := FindKthTail(head, -1)
	if existNeg1 {
		t.Logf("the reciprocal 1th node: %d\n", kthNeg1.Value)
	} else {
		t.Errorf("the reciprocal %dth node does not exist.\n", -1)
	}

	kth0, exist1 := FindKthTail(head, 0)
	if exist1 {
		t.Logf("the reciprocal 1th node: %d\n", kth0.Value)
	} else {
		t.Errorf("the reciprocal %dth node does not exist.\n", 0)
	}

	kth1, exist1 := FindKthTail(head, k)
	if exist1 {
		k++
		t.Logf("the reciprocal 1th node: %d\n", kth1.Value)
	} else {
		t.Errorf("the reciprocal %dth node does not exist.\n", k)
	}

	kth2, exist2 := FindKthTail(head, k)
	if exist2 {
		k++
		t.Logf("the reciprocal 2th node: %d\n", kth2.Value)
	} else {
		t.Errorf("the reciprocal %dth node does not exist.\n", k)
	}

	kth3, exist3 := FindKthTail(head, k)
	if exist3 {
		k++
		t.Logf("the reciprocal 3th node: %d\n", kth3.Value)
	} else {
		t.Errorf("the reciprocal %dth node does not exist.\n", k)
	}

	kth4, exist4 := FindKthTail(head, k)
	if exist4 {
		k++
		t.Logf("the reciprocal 4th node: %d\n", kth4.Value)
	} else {
		t.Errorf("the reciprocal %dth node does not exist.\n", k)
	}

	kth5, exist5 := FindKthTail(head, k)
	if exist5 {
		k++
		t.Logf("the reciprocal 4th node: %d\n", kth5.Value)
	} else {
		t.Errorf("the reciprocal %dth node does not exist.\n", k)
	}
}
