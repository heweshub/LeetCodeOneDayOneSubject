package main

import "sort"

func smallestRangeI(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	sort.Ints(nums)
	if nums[0]+k >= nums[n-1]-k {
		return 0
	} else {
		return nums[n-1] - k - nums[0] - k
	}
}
