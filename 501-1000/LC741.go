package main

import "math"

func cherryPickup(grid [][]int) int {
	n := len(grid)
	f := make([][][]int, n*2-1)
	for i := range f {
		f[i] = make([][]int, n)
		for j := range f[i] {
			f[i][j] = make([]int, n)
			for k := range f[i][j] {
				f[i][j][k] = math.MinInt32
			}
		}
	}
	f[0][0][0] = grid[0][0]
	for k := 1; k < 2*n-1; k++ {
		for x1 := max(k-n+1, 0); x1 <= min(n-1, k); x1++ {
			y1 := k - x1
			if grid[x1][y1] == -1 {
				continue
			}
			for x2 := x1; x2 <= min(k, n-1); x2++ {
				y2 := k - x2
				if grid[x2][y2] == -1 {
					continue
				}
				res := f[k-1][x1][x2]
				if x1 > 0 {
					res = max(res, f[k-1][x1-1][x2])
				}
				if x2 > 0 {
					res = max(res, f[k-1][x1][x2-1])
				}
				if x1 > 0 && x2 > 0 {
					res = max(res, f[k-1][x1-1][x2-1])
				}
				res += grid[x1][y1]
				if x2 != x1 {
					res += grid[x2][y2]
				}
				f[k][x1][x2] = res
			}
		}
	}
	return max(f[n*2-2][n-1][n-1], 0)
}
