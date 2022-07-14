package main

import "strconv"

func translate(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 1; i < n; i++ {
		dp[i+1] = dp[i]
		if a, _ := strconv.Atoi(s[i-1 : i+1]); a >= 10 && a <= 25 {
			dp[i+1] = dp[i-1] + dp[i]
		}
	}
	return dp[n]
}
