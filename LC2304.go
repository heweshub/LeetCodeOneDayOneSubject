package main

import "math"

func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
	pre := grid[0]
	f := make([]int, n)
	for i := 1; i < m; i++ {
		for j, g := range grid[i] {
			f[j] = math.MaxInt32
			for k, v := range grid[i-1] {
				f[j] = min(f[j], pre[k]+moveCost[v][j])
			}
			f[j] += g
		}
		pre, f = f, pre
	}
	ans := math.MaxInt32
	for _, v := range pre {
		ans = min(ans, v)
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
