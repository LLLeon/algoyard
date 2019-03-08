package binary

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	value := 3
	wantedIndex := 2

	i := BinarySearch(list, value)

	if i == wantedIndex {
		t.Log("Find the value, index:", i)
	} else if i == -1 {
		t.Log("There is no number or value in the array.")
	} else {
		t.Fatalf("Wanted index: %d, get index: %d\n", wantedIndex, i)
	}
}
