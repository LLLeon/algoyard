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
	q := NewCircularQueue(10)

	// When the queue is full, no data is stored at the tail location,
	// so the circular queue wastes one storage space.
	for i := 0; i < 9; i++ {
		if ok := q.Enqueue(i); !ok {
			t.Fatal("CircularQueue enqueue error.")
		}
	}

	for i := 0; i < 9; i++ {
		item, ok := q.Dequeue()
		if !ok {
			t.Fatal("CircularQueue dequeue error")
		}
		fmt.Printf("CircularQueue dequeue item: %d\n", item)
	}
}
