package main

func movingCount(m int, n int, k int) int {
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, vis)
}

func dfs(i, j, m, n, k int, vis [][]bool) int {
	if i < 0 || i >= m || j < 0 || j >= n ||
		countNum(i)+countNum(j) > k || vis[i][j] {
		return 0
	}
	vis[i][j] = true
	return dfs(i+1, j, m, n, k, vis) + dfs(i-1, j, m, n, k, vis) +
		dfs(i, j+1, m, n, k, vis) + dfs(i, j-1, m, n, k, vis) + 1
}

func countNum(x int) (ans int) {
	for x > 0 {
		gewei := x % 10
		ans += gewei
		x /= 10
	}
	return
}
