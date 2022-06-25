package main

import "sort"

func findPairs(nums []int, k int) (ans int) {
	sort.Ints(nums)
	y, n := 0, len(nums)
	for x, num := range nums {
		if x == 0 || num != nums[x-1] {
			for y < n && (nums[y] < num+k || y <= x) {
				y++
			}
			if y < n && nums[y] == num+k {
				ans++
			}
		}
	}
	return
}
