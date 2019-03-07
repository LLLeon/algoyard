package quick

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

	i := partition(list, start, end)
	separateSort(list, start, i-1)
	separateSort(list, i+1, end)
}

func partition(list []int, start, end int) int {
	pivot := list[end]
	i := start

	for j := start; j < end; j++ {
		if list[j] < pivot {
			if i != j {
				list[i], list[j] = list[j], list[i]
			}
			i++
		}
	}

	list[i], list[end] = list[end], list[i]

	return i
}
