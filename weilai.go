package main

import "fmt"

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	nums := make([]int, n)
// 	for i := range nums {
// 		fmt.Scan(&nums[i])
// 	}
// 	// fmt.Println(nums)
// 	dp := make([]int, n)
// 	dp[0] = nums[0]
// 	for i := 1; i < n; i++ {
// 		dp[i] = max(dp[i]+dp[i-1], dp[i])
// 	}
// 	fmt.Println(dp[n-1])
// }

// func max(x, y int) int {
// 	if x < y {
// 		return y
// 	}
// 	return x
// }

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	nums := make([]int, n)
// 	for i := range nums {
// 		fmt.Scan(&nums[i])
// 	}
// 	idxs := make([]int, n)
// 	for i := range idxs {
// 		fmt.Scan(&idxs[i])
// 	}
// 	fmt.Println(nums)
// 	fmt.Println(idxs)
// }
var directions = [][]int{{1, 0}, {0, 1}}
var sum = 0

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	nums := make([][]int, n)
	for i := range nums {
		nums[i] = make([]int, m)
	}
	for i := range nums {
		for j := range nums[i] {
			fmt.Scan(&nums[i][j])
		}
	}
	// fmt.Println(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	dp[0][0] = nums[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (i == 0 || i == n-1) && j > 0 {
				dp[i][j] = dp[i][j-1] + nums[i][j]
			} else if (j == 0 || j == m-1) && i > 0 {
				dp[i][j] = dp[i-1][j] + nums[i][j]
			} else {
				dp[i][j] = nums[i][j] + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp[n-1][m-1])
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
