package merge

// MergeSort implements merge sort.
// Time complexity: O(n*log(n)).
// Space complexity: O(n).
func MergeSort(list []int, n int) {
	if n <= 1 {
		return
	}

	mergeSort(list, 0, n-1)
}

func mergeSort(list []int, start, end int) {
	// recursive termination condition
	if start >= end {
		return
	}

	mid := (start + end) / 2

	// divide and conquer recursive
	mergeSort(list, start, mid)
	mergeSort(list, mid+1, end)
	merge(list, start, mid, end)
}

func merge(list []int, start, mid, end int) {
	tmpList := make([]int, end-start+1)

	// i and j point to the first number of list[start:mid+1]
	// and list[mid+1:end+1], respectively.
	// k is the index in tmpList.
	i := start
	j := mid + 1
	k := 0

	// compares the numbers in two subarrays in turn,
	// placing the smaller number in tmpList.
	for ; i <= mid && j <= end; k++ {
		if list[i] <= list[j] {
			tmpList[k] = list[i]
			i++
		} else {
			tmpList[k] = list[j]
			j++
		}
	}

	// place the remaining numbers of the two subarrays
	// into tmpList in turn (if any).
	for ; i <= mid; i++ {
		tmpList[k] = list[i]
		k++
	}

	for ; j <= end; j++ {
		tmpList[k] = list[j]
		k++
	}

	// copies the sorted numbers from tmpList to the original
	// position of the original list.
	copy(list[start:end+1], tmpList)
}
