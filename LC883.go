package main

import (
	"fmt"
	"math"
)

func projectionArea(grid [][]int) int {

	sum := 0
	// 顶部看到的影子
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 0 {
				sum += 1
			}
		}
	}
	// 前面看到的影子
	for i := range grid {
		maxItem := math.MinInt32
		for j := range grid[i] {
			if grid[i][j] > maxItem {
				maxItem = grid[i][j]
			}
		}
		sum += maxItem
	}
	// 侧面看到的影子
	for j := 0; j < len(grid[0]); j++ {
		maxItem := math.MinInt32
		for i := 0; i < len(grid); i++ {
			if grid[i][j] > maxItem {
				maxItem = grid[i][j]
			}
		}
		sum += maxItem
	}
	return sum
}

func main() {
	//fmt.Println(projectionArea([][]int{{1, 2}, {3, 4}}))
	fmt.Println(projectionArea([][]int{{2}}))
}
