package main

import "sort"

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right {
			right = max(right, intervals[i][1])
		} else {
			ans = append(ans, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		}
	}
	ans = append(ans, []int{left, right})
	return
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
