package main

import "sort"

func maximumGroups(grades []int) (ans int) {
	sort.Ints(grades)
	for i, preSum, k, n := 0, 0, 0, len(grades); i+k < n; ans++ {
		i0, sum := i, 0
		for ; i < n && (sum <= preSum || i <= i0+k); i++ {
			sum += grades[i]
		}
		if sum <= preSum {
			break
		}
		preSum, k = sum, i-i0
	}
	return
}
