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
		value := list[i]
		j := i - 1

		fmt.Printf("enter j loop, value=%d\n", value)
		// find the position to insert
		for ; j >= 0; j-- {
			if list[j] > value {
				// data movement
				list[j+1] = list[j]
			} else {
				break
			}

			fmt.Printf("i=%d, j=%d, list=%v\n", i, j, list)
		}

		// insert data
		list[j+1] = value

		fmt.Printf("exit j loop, list[%d]=value=%d\n", j+1, list[j+1])
		fmt.Printf("now list=%v\n", list)
		fmt.Printf("\n")
	}

	return list
}
