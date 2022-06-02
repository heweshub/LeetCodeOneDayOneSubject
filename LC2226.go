package main

import "sort"

func maximumCandies(candies []int, k int64) int {
	total := 0
	for _, v := range candies {
		total += v
	}
	if total < int(k) {
		return 0
	}
	return sort.Search(total+1, func(x int) bool {
		if x == 0 {
			return false
		}
		var cnt int
		for _, v := range candies {
			cnt += v / int(k)
		}
		return cnt < int(k)
	})
}
