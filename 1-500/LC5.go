package main

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	max_len, begin := 1, 0
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for l := 2; l < n+1; l++ {
		// 左边界
		for i := 0; i < n; i++ {
			// 右边界
			r := l + i - 1
			if r >= n {
				break
			}
			if s[i] != s[r] {
				dp[i][r] = false
			} else {
				if r-i < 3 {
					dp[i][r] = true
				} else {
					dp[i][r] = dp[i+1][r-1]
				}
			}
			if dp[i][r] && r-i+1 > max_len {
				max_len = r - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+max_len]
}

func longestPalindrome1(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < n; i++ {
		left1, right1 := expand(s, i, i)
		left2, right2 := expand(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && len(s) > right {
		if s[left] == s[right] {
			left--
			right++
		}
	}
	return left + 1, right - 1
}
