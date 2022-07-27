package main

func canJump(nums []int) bool {
	n, right := len(nums), 0
	for i := 0; i < n; i++ {
		if i <= right {
			right = max(right, i+nums[i])
			if right >= n-1 {
				return true
			}
		}
	}
	return false
}
