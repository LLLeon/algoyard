package quick

import "testing"

func TestQuickSort(t *testing.T) {
	list := []int{5, 7, 3, 8, 4, 6}
	t.Log("Original list:", list)
	QuickSort(list, len(list))
	t.Log("Sorting result:", list)
}

func TestQSort(t *testing.T) {
	list := []int{5, 7, 3, 8, 4, 6}
	t.Log("Original list:", list)
	QSort(list, len(list))
	t.Log("Sorting result:", list)
}
