package main

func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l <= r {
		mid := (r + l) >> 1
		cur := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cur++
			}
		}
		if cur <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}
	return ans
}
