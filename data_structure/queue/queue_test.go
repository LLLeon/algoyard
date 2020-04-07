package queue

import (
	"fmt"
	"testing"
)

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(10)

	for i := 0; i < 10; i++ {
		if ok := q.Enqueue(i); !ok {
			t.Fatal("ArrayQueue enqueue error.")
		}
	}

	for i := 0; i < 10; i++ {
		item, ok := q.Dequeue()
		if !ok {
			t.Fatal("ArrayQueue dequeue error")
		}
		fmt.Printf("ArrayQueue dequeue item: %d\n", item)
	}
}

func TestCircularQueue(t *testing.T) {
	q := NewCircularQueue(3)

	// should be successful
	if !q.Enqueue(1) {
		t.Error("enqueue 1 failed")
	}
	if !q.Enqueue(2) {
		t.Error("enqueue 2 failed")
	}
	if !q.Enqueue(3) {
		t.Error("enqueue 3 failed")
	}

	// should fail
	if q.Enqueue(4) {
		t.Error("enqueue 4 success")
	}

	t.Logf("q ->%v\n", q)

	// should be 1
	if q.Front() != 1 {
		t.Error("the front item is not 1")
	}
	if !q.Dequeue() {
		t.Error("dequeue failed")
	}

	// should be successful
	if !q.Enqueue(4) {
		t.Error("enqueue 4 failed")
	}
	if q.Rear() != 4 {
		t.Error("the last item is not 4")
	}

	t.Logf("q ->%v\n", q)

	t.Log("success")
}
