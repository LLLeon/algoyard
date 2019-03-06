package selection

import "fmt"

// SelectionSort implements selection sort.
// The best situation: [1, 2, 3, 4, 5, 6], time complexity: O(n^2).
// The worst situation: [6, 5, 4, 3, 2, 1], time complexity: O(n^2).
// Average time complexity: O(n^2).
// Space complexity: O(1).
func SelectionSort(list []int, n int) []int {
	if n <= 1 {
		return list
	}

	for i := 0; i < n; i++ {
		minIndex := i

		fmt.Printf("enter j loop\n")
		for j := i + 1; j < n; j++ {
			fmt.Printf("i=%d, j=%d, minIndex=%d, list=%v\n", i, j, minIndex, list)

			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		fmt.Printf("exit j loop\n")
		fmt.Printf("exchange data: list[i]: list[%d]=%d <-> list[minIndex]: list[%d]=%d\n",
			i, list[i], minIndex, list[minIndex])

		list[i], list[minIndex] = list[minIndex], list[i]

		fmt.Printf("after exchanging data: %v\n", list)
		fmt.Printf("\n")
	}

	return list
}
