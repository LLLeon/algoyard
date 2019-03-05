package bubble

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	oldList := []int{6, 5, 4, 3, 2, 1}
	t.Log("Original list:", oldList)
	newList := BubbleSort(oldList, len(oldList))
	t.Log("Sorting result:", newList)
}
