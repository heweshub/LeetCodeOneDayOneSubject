package main

import "sort"

func eraseOverlapIntervals(inter [][]int) int {
	n := len(inter)
	if n == 0 {
		return 0
	}
	sort.Slice(inter, func(i, j int) bool {
		return inter[i][1] < inter[j][1]
	})
	ans, right := 1, inter[0][1]
	for _, p := range inter[1:] {
		// 表明是两个不重叠的区间
		if p[0] >= right {
			ans++
			right = p[1]
		}
	}
	return n - ans
}
