package main

import "sort"

const inf int = 0x3f3f3f

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = inf
		for _, c := range coins {
			// coin大于amount
			if c > i {
				break
			}
			// 状态转移方程
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}
	if dp[amount] == inf {
		dp[amount] = -1
	}
	return dp[amount]
}
