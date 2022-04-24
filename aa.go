package main

import (
	"fmt"
	"math"
)

func countLatticePoints(circles [][]int) (ans int) {
	s, x, z, y := math.MinInt32, math.MaxInt32, math.MaxInt32, math.MinInt32
	n := len(circles)
	for i := 0; i < n; i++ {
		if circles[i][0]-circles[i][2] < z {
			z = circles[i][0] - circles[i][2]
		}
		if circles[i][0]+circles[i][2] > y {
			y = circles[i][0] + circles[i][2]
		}
		if circles[i][1]-circles[i][2] < x {
			x = circles[i][1] - circles[i][2]
		}
		if circles[i][1]+circles[i][2] > s {
			s = circles[i][1] + circles[i][2]
		}
	}
	vis := make([][]bool, s+1)
	for i := 0; i < s+1; i++ {
		vis[i] = make([]bool, y+1)
	}
	for k := 0; k < n; k++ {
		yy1 := circles[k][1] - circles[k][2]
		yy2 := circles[k][1] + circles[k][2]
		for i := yy1; i <= yy2; i++ {
			xx1 := circles[k][0] - circles[k][2]
			xx2 := circles[k][0] + circles[k][2]
			for j := xx1; j <= xx2; j++ {
				//fmt.Println(i, j, isInCircles(i, j, circles[k]))
				if isInCircles(j, i, circles[k]) {
					if !vis[i][j] {
						ans++
						vis[i][j] = true
					}
				}
			}
		}
	}
	return
}

func isInCircles(x, y int, circles []int) bool {
	if isIn(x, y, circles[0], circles[1], circles[2]) {
		return true
	}
	return false
}
func isIn(x, y, qx, qy, r int) bool {
	var length float64 = math.Sqrt(float64(abs1(x-qx)*abs1(x-qx) + abs1(y-qy)*abs1(y-qy)))
	if length <= float64(r) {
		return true
	}
	return false
}

func abs1(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	//fmt.Println(countLatticePoints([][]int{{2, 2, 1}}))
	//fmt.Println(countLatticePoints([][]int{{2, 2, 2}, {3, 4, 1}}))
	fmt.Println(countLatticePoints([][]int{{10, 10, 8}, {5, 9, 2}, {7, 1, 1}}))
}
