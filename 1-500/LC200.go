package main

func numIslands(grid [][]byte) (ans int) {
	n, m := len(grid), len(grid[0])
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == '0' || vis[i][j] {
			return
		}
		vis[i][j] = true
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i+1, j)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && vis[i][j] {
				ans++
				dfs(i, j)
			}
		}
	}
	return
}
