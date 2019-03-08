package binary

// BinarySearch implements binary search.
// The working scenario for this implementation
// is an ordered array with no duplicate numbers.
// Time complexity: O(logN).
func BinarySearch(list []int, value int) int {
	var (
		n    = len(list)
		low  = 0
		high = n - 1
		mid  = 0
	)

	if n == 0 {
		return -1
	}

	for low <= high {
		// prevent overflow, can be written as:
		// mid = low + ((high - low) >> 1)
		mid = (low + high) / 2

		if list[mid] == value {
			return mid
		} else if list[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
