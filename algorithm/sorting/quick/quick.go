package quick

import "fmt"

// QuickSort implements quick sort.
//
// Depending on whether the partition is balanced or not,
// the best situation: time complexity: O(n*log(n)).
// The worst situation: time complexity: O(n^2).
// Average time complexity: O(n*log(n)).
//
// The best situation, space complexity: O(log(n)).
// The worst situation, space complexity: O(n).
func QuickSort(list []int, n int) {
	if n <= 1 {
		return
	}

	separateSort(list, 0, n-1)
}

func separateSort(list []int, start, end int) {
	if start >= end {
		return
	}

	pivotIndex := partition(list, start, end)
	separateSort(list, start, pivotIndex-1)
	separateSort(list, pivotIndex+1, end)
}

func partition(list []int, start, end int) int {
	pivot := list[end]
	i := start

	fmt.Printf("original list: %v\n", list)
	fmt.Printf("pivot: %d\n", pivot)

	for j := start; j < end; j++ {
		fmt.Printf("enter loop, i = %d, j = %d\n", i, j)
		if list[j] < pivot {
			if i != j {
				list[i], list[j] = list[j], list[i]
				fmt.Printf("exchanged list[%d] <---> list[%d]\n", i, j)
			}
			i++
		}
		fmt.Printf("exit loop, i = %d, j = %d\n", i, j)
		fmt.Printf("\n")
	}

	// exchange the value of i with pivot
	list[i], list[end] = list[end], list[i]

	fmt.Printf("list[%d] <---> list[%d]\n", i, end)
	fmt.Printf("changed list: %v\n", list)
	fmt.Printf("\n")

	// returns the index of pivot
	return i
}

// QSort is another implementation of quick sort.
// This implementation is more intuitive.
// Compared with this implementation, the above is
// simply too ugly.
func QSort(list []int, n int) {
	if n <= 1 {
		return
	}

	pivot := list[0]
	head, tail := 0, n-1

	for i := 1; i <= tail; {
		if list[i] > pivot {
			list[i], list[tail] = list[tail], list[i]
			tail--
		} else {
			list[i], list[head] = list[head], list[i]
			head++
			i++
		}
	}

	frontPart := list[:head]
	latterPart := list[head+1:]

	QSort(frontPart, len(frontPart))
	QSort(latterPart, len(latterPart))
}
