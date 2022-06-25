package main

import "sort"

func checkPossibility(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			nums[i] = y
			if sort.IntsAreSorted(nums) {
				return true
			}
			nums[i] = x
			nums[i+1] = x
			return sort.IntsAreSorted(nums)
		}
	}
	return true
}
