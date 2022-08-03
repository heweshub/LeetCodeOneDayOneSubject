package main

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}
	i2, i3, i5 := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		next2, next3, next5 := dp[i2]*2, dp[i3]*3, dp[i5]*5
		dp[i] = min(next2, min(next3, next5))
		if dp[i] == next2 {
			i2++
		} else if dp[i] == next3 {
			i3++
		} else if dp[i] == next5 {
			i5++
		}
	}
	return dp[n-1]
}

// func min(x, y int) int {
// 	if x < y {
// 		return x
// 	}
// 	return y
// }
