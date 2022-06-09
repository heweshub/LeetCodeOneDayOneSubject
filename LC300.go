package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := []int{}
	for i := 0; i < len(nums); i++ {
		dp = append(dp, 1)
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return maxOfArray(dp)
}
