package main

import "fmt"

func maxConsecutiveAnswers(a string, k int) int {
	n := len(a)
	dp := make([]int, n)
	cnt := 1
	dp[0] = 1
	for i := 1; i < n; i++ {
		if a[i-1] == a[i] {
			cnt++
			dp[i] = cnt
		} else {
			cnt = 1
			dp[i] = 1
		}
	}
	fmt.Println(dp)
	return max(dp) + k
}

func main() {
	fmt.Println(maxConsecutiveAnswers("TTFF", 2))
}

func max(a []int) (ans int) {
	for i := 0; i < len(a); i++ {
		if ans < a[i] {
			ans = a[i]
		}
	}
	return
}
