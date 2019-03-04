package queue

import (
	"testing"
	"fmt"
)

func TestArrayQueue(t *testing.T) {
	q := NewArrayQueue(10)

	for i := 0; i < 10; i++ {
		if ok := q.Enqueue(i); !ok {
			t.Error("enqueue error.")
		}
	}

	for i := 0; i < 10; i++ {
		item, ok := q.Dequeue()
		if !ok {
			t.Error("dequeue error")
		}
		fmt.Printf("dequeue item: %d\n", item)
	}
}
