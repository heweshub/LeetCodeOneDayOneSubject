package main

var coordinate = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func closedIsland(grid [][]int) int {
	flag := make([][]bool, len(grid))
	for i := range flag {
		flag[i] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		dfss(flag, 0, i, grid)
		dfss(flag, len(grid)-1, i, grid)
	}
	for i := 0; i < len(grid); i++ {
		dfss(flag, i, 0, grid)
		dfss(flag, i, len(grid[0])-1, grid)
	}
	var ans int = 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 0 {
				ans++
				dfss(flag, i, j, grid)
			}
		}
	}
	return ans
}

func dfss(flag [][]bool, r, c int, grid [][]int) {
	if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || flag[r][c] {
		return
	}
	flag[r][c] = true
	if grid[r][c] == 1 {
		return
	}
	grid[r][c] = 1
	for _, value := range coordinate {
		dfss(flag, r+value[0], c+value[1], grid)
	}
}
