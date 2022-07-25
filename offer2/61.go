package main

func isStraight(nums []int) bool {
	m := make(map[int]bool)
	l, h := 0, 14
	for _, v := range nums {
		if v == 0 {
			continue
		}
		l = max(l, v)
		h = min(h, v)
		if _, ok := m[v]; ok {
			return false
		} else {
			m[v] = true
		}
	}
	return true
}

// func max(x, y int) int {
// 	if x < y {
// 		return y
// 	}
// 	return x
// }

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
