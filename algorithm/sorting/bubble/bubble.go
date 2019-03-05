package bubble

func BubbleSort(list []int, n int) []int {
	if n <= 1 {
		return list
	}

	for i := 0; i < n; i++ {
		isExchange := false

		for j := 0; j < n-i-1; j++ {
			if list[j] > list[j+1] {
				tmp := list[j]
				list[j] = list[j+1]
				list[j+1] = tmp
				isExchange = true
			}
		}

		if !isExchange {
			break
		}
	}

	return list
}
