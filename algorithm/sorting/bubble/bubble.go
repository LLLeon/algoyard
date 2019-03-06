package bubble

import "fmt"

// BubbleSort implements bubble sort.
// The best situation: [1, 2, 3, 4, 5, 6], just one bubble,
// time complexity: O(n).
// The worst situation: [6, 5, 4, 3, 2, 1], it takes six bubbles,
// time complexity: O(n^2).
// Average time complexity: O(n^2).
// Space complexity: O(1).
func BubbleSort(list []int, n int) []int {
	if n <= 1 {
		return list
	}

	for i := 0; i < n; i++ {
		// Used to exit the loop early.
		isExchange := false

		fmt.Printf("enter j loop\n")
		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				isExchange = true

				fmt.Printf("i=%d, j=%d, list=%v\n", i, j, list)
			}
		}
		fmt.Printf("exit j loop\n")
		fmt.Printf("\n")

		if !isExchange {
			break
		}
	}

	return list
}
