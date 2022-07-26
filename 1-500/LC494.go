package main

func findTargetSumWays(nums []int, target int) (count int) {
	var backtrack func(int, int)
	backtrack = func(idx, sum int) {
		if idx == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		backtrack(idx+1, sum+nums[idx])
		backtrack(idx+1, sum-nums[idx])
	}
	backtrack(0, 0)
	return
}
