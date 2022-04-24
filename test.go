package main

import (
	"fmt"
)

var vis [][]int
var matrix [][]int
var ans int

func dfs(n, m, i, j int) int {
	if i < 0 || j < 0 || i >= n || j >= m || matrix[i][j] == 1 {
		return ans
	} else if vis[i][j] != 1 && matrix[i][j] == 0 {
		vis[i][j] = 1
		ans++
		dfs(n, m, i+1, j)
		dfs(n, m, i-1, j)
		dfs(n, m, i, j+1)
		dfs(n, m, i, j-1)
	}
	return ans
}

//4
//7
//1 0 0 1 0 0 0
//1 0 0 0 0 1 1
//0 0 0 1 0 0 0
//1 1 0 1 1 0 0
//4
//7
//1 0 0 1 0 0 0
//1 0 0 1 0 1 1
//0 0 0 1 0 0 0
//1 1 0 1 1 0 0

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	matrix = make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	//n, m = 4, 7
	//matrix = [][]int{{1, 0, 0, 1, 0, 0, 0}, {1, 0, 0, 0, 0, 1, 1}, {0, 0, 0, 1, 0, 0, 0}, {1, 1, 0, 1, 1, 0, 0}}
	//matrix = [][]int{{1, 0, 0, 1, 0, 0, 0}, {1, 0, 0, 1, 0, 1, 1}, {0, 0, 0, 1, 0, 0, 0}, {1, 1, 0, 1, 1, 0, 0}}
	vis = make([][]int, n)
	for i := range vis {
		vis[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			vis[i][j] = matrix[i][j]
		}
	}
	maxCount := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ans = 0
			count := dfs(n, m, i, j)
			if count > maxCount {
				maxCount = count
			}
		}
	}
	fmt.Println(maxCount)
}
