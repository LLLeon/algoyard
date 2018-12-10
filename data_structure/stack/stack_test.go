package stack

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	s := NewArrayStack(10)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for _, item := range items {
		pushed := s.Push(item)
		if pushed {
			fmt.Printf("[push] push %s success\n", item)
		} else {
			fmt.Printf("[push] failed to push %s\n", item)
			return
		}
	}

	fmt.Printf("\n")

	for i := s.count - 1; i >= 0; i-- {
		want := items[i]
		get, ok := s.Pop()
		if !ok {
			fmt.Println("[pop] failed to pop")
			return
		}

		if get != want {
			fmt.Println("[pop] the element out of the stack is not correct")
			return
		}

		fmt.Printf("[pop] want: %s, get: %s, elements match correctly\n", want, get)
	}

	fmt.Println("both push and pop operations were successful")
}

func TestDynamicScalingStack(t *testing.T) {
	s := NewArrayStack(10)
	items := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l"}

	for _, item := range items {
		pushed := s.Push(item)
		if pushed {
			fmt.Printf("[push] push %s success\n", item)
			fmt.Printf("[push] stack size: %d\n", s.size)
			fmt.Printf("[push] number of elements in the stack: %d\n", s.count)
			fmt.Printf("\n")
		} else {
			fmt.Printf("[push] failed to push %s\n", item)
			return
		}
	}

	wantSize := int64(20)

	if s.size == wantSize {
		fmt.Printf("stack dynamic scaling successfully, want: %d, get: %d\n", wantSize, s.size)
	} else {
		fmt.Printf("stack dynamic scaling failed, want: %d, get: %d\n", wantSize, s.size)
	}

	fmt.Printf("\n")

	for i := s.count - 1; i >= 0; i-- {
		want := items[i]
		get, ok := s.Pop()
		if !ok {
			fmt.Println("[pop] failed to pop")
			return
		}

		if get != want {
			fmt.Println("[pop] the element out of the stack is not correct")
			return
		}

		fmt.Printf("[pop] want: %s, get: %s, elements match correctly\n", want, get)
	}

	fmt.Println("both push and pop operations were successful")
}
