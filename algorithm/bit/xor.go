package bit

// use map
// Time complexity O(n), Space complexity O(n)
func DuplicateIntI(l []int) []int {
	if len(l) == 0 {
		return []int{}
	}

	res := make([]int, 0)
	// k: v -> int: count
	m := make(map[int]int)

	for _, v := range l {
		if c, ok := m[v]; ok {
			m[v] = c + 1
		} else {
			m[v] = 1
		}
	}

	for i, c := range m {
		if c == 1 {
			res = append(res, i)
		}
	}

	return res
}

func DuplicateIntII(l []int) []int {
	if len(l) == 0 {
		return []int{}
	}

	// 用于将所有的数异或起来
	k := 0

	for _, num := range l {
		k ^= num
	}

	// 获得k中最低位的1
	mask := 1

	// mask = k & (-k) 这种方法也可以得到 mask
	for k&mask == 0 {
		mask <<= 1
	}

	a, b := 0, 0

	for _, num := range l {
		if num&mask == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	res := []int{a, b}
	return res
}
