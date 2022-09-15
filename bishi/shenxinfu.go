package main

import (
	"fmt"
)

func main() {
	// var n, x int
	// fmt.Scan(&n, &x)
	// a := make([]int, n)
	// for i := range a {
	// 	fmt.Scan(&a[i])
	// }
	// mmin := max(0, a[0]-x)
	// mmax := a[0] + x
	// cnt := 0
	// for i := 1; i < n; i++ {
	// 	tmpmax := a[i] + x
	// 	tmpmin := max(0, a[i]-x)
	// 	mmax = min(mmax, tmpmax)
	// 	mmin = max(mmin, tmpmin)
	// 	if mmin > mmax {
	// 		cnt++
	// 		mmin = tmpmin
	// 		mmax = tmpmax
	// 	}
	// }
	// fmt.Print(cnt)
	fmt.Println(minDistance("horse", "rose"))
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minDistance(a, b string) int {
	na, nb := len(a), len(b)
	dp := make([][]int, na)
	for i := range dp {
		dp[i] = make([]int, nb)
	}
	for i := 0; i < na; i++ {
		dp[i][0] = i
	}
	for i := 0; i < nb; i++ {
		dp[0][i] = i
	}
	for i := 1; i < na; i++ {
		for j := 1; j < nb; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j-1]+1, dp[i-1][j]+1), dp[i][j-1]+1)
			}
		}
	}
	fmt.Println(dp)
	return dp[na-1][nb-1]
}
