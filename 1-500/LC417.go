package main

var direct = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func pacificAtlantic(h [][]int) (ans [][]int) {
	m, n := len(h), len(h[0])
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range h {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}
	var dfs func(i, j int, ocean [][]bool)
	dfs = func(i, j int, ocean [][]bool) {
		if ocean[i][j] {
			return
		}
		ocean[i][j] = true
		for k := range direct {
			if nx, ny := i+direct[k][0], j+direct[k][1]; nx >= 0 && nx < m && ny >= 0 && ny < n && h[i][j] <= h[nx][ny] {
				dfs(nx, ny, ocean)
			}
		}
	}
	for j := 0; j < n; j++ {
		dfs(0, j, pacific)
	}
	for i := 1; i < m; i++ {
		dfs(i, 0, pacific)
	}
	for j := 0; j < n; j++ {
		dfs(m-1, j, atlantic)
	}

	for i := 0; i < m; i++ {
		dfs(i, n-1, atlantic)
	}

	for i := range pacific {
		for j := range pacific[i] {
			if pacific[i][j] && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}
