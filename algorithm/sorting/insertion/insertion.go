package insertion

import "fmt"

// InsertionSort implements insertion sort.
// This implementation finds where to insert from the end of the
// sorted part to the head.
// The best situation: [1, 2, 3, 4, 5, 6], time complexity: O(n).
// The worst situation: [6, 5, 4, 3, 2, 1], time complexity: O(n^2).
// Average time complexity: O(n^2).
// Space complexity: O(1).
func InsertionSort(list []int, n int) []int {
	if n <= 1 {
		return list
	}

	for i := 1; i < n; i++ {
		// the value to insert
		value := list[i]
		j := i - 1

		fmt.Printf("the value to insert: list[%d] = %d\n", i, value)

		// find the position to insert
		for ; j >= 0; j-- {
			if list[j] > value {
				// data movement
				list[j+1] = list[j]
			} else {
				break
			}

			fmt.Printf("i = %d, j = %d\n", i, j)
			fmt.Printf("after moving list[%d], list ---> %v\n", j, list)
		}

		// insert data
		list[j+1] = value

		fmt.Printf("the position to insert: list[%d]\n", j+1)
		fmt.Printf("now list ---> %v\n", list)
		fmt.Printf("\n")
	}

	return list
}
