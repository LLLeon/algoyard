package merge

import "testing"

func TestMergeSort(t *testing.T) {
	list := []int{6, 5, 4, 3, 2, 1}
	t.Log("Original list:", list)
	MergeSort(list, len(list))
	t.Log("Sorting result:", list)
}
