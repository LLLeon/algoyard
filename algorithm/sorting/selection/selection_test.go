package selection

import "testing"

func TestSelectionSort(t *testing.T) {
	oldList := []int{6, 5, 4, 3, 2, 1}
	t.Log("Original list:", oldList)
	newList := SelectionSort(oldList, len(oldList))
	t.Log("Sorting result:", newList)
}
