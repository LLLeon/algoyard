package insertion

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	oldList := []int{6, 4, 5, 3, 2, 1}
	t.Log("Original list:", oldList)
	newList := InsertionSort(oldList, len(oldList))
	t.Log("Sorting result:", newList)
}
