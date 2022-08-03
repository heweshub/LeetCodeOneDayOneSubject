package main

func maximalSquare(m [][]byte) int {
	dp := make([][]int, len(m))
	maxSize := 0
	for i := 0; i < len(m); i++ {
		dp[i] = make([]int, len(m[i]))
		for j := 0; j < len(m[i]); j++ {
			dp[i][j] = int(m[i][j] - '0')
			if dp[i][j] == 1 {
				maxSize = 1
			}
		}
	}
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				if dp[i][j] > maxSize {
					maxSize = dp[i][j]
				}
			}
		}
	}
	return maxSize * maxSize
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
