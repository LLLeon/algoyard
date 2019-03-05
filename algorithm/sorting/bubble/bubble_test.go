package bubble

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	oldList := []int{4, 5, 6, 1, 2, 3, 3, 2}
	t.Log("Original list:", oldList)
	newList := BubbleSort(oldList, len(oldList))
	t.Log("Sorting result:", newList)
}
