package main

import (
	"fmt"
)

//func getMinimumTime(time []int, fruits [][]int, limit int) int {
//	times := 0
//	for _, f := range fruits {
//		t := 0
//		//fmt.Println(f[1] % limit)
//		if f[1]%limit == 0 {
//			t = time[f[0]] * (f[1] / limit)
//		} else {
//			t = time[f[0]] * (f[1]/limit + 1)
//		}
//		times += t
//	}
//	//fmt.Println(times)
//	return times
//}
//func main() {
//	time := []int{1}
//	fruits := [][]int{{0, 3}, {0, 5}}
//	limit := 3
//	getMinimumTime(time, fruits, limit)
//
//}

func conveyorBelt(matrix []string, start []int, end []int) int {
	n, m := len(matrix), len(matrix[0])
	startX, startY := start[0], start[1]
	endX, endY := end[0], end[1]
	var win bool
	for i := 0; i < n*m; i++ {
		tag := matrix[startX][startY]
		if tag == '>' {
			startY++
		} else if tag == '<' {
			startY--
		} else if tag == '^' {
			startX--
		} else {
			startX++
		}
		if startX == endX && startY == endY {
			win = true
			break
		}
		if startX < 0 || startX >= n || startY < 0 || startY >= m {
			win = false
			break
		}
	}
	if win {
		return 0
	}
	// 不能直接走通的
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	bb := make([][]byte, n)
	for i := range bb {
		bb[i] = make([]byte, m)
	}
	for i := range matrix {
		for j := range matrix[i] {
			bb[i][j] = matrix[i][j]
		}
	}
	startX, startY = start[0], start[1]
	var dfs func(s, e, depth int) int
	dfs = func(s, e, depth int) int {
		if s == end[0] && e == end[1] {
			return depth
		}
		vis[s][e] = true
		tag := bb[s][e]
		if tag == '>' && !vis[s][e+1] && e+1 < m {
			startY++
		} else if tag == '<' && !vis[s][e-1] && e-1 >= 0 {
			startY--
		} else if tag == '^' && !vis[s-1][e] && s-1 >= 0 {
			startX--
		} else if tag == 'v' && !vis[s+1][e] && s+1 < n {
			startX++
		} else {
			for _, v := range []byte{'>', '<', '^', 'v'} {
				if tag == v {
					continue
				}
				bb[s][e] = v
				return dfs(s, e, depth+1)
			}
		}
		return dfs(s, e, depth)
	}
	return dfs(startX, startY, 0)
}

func main() {
	fmt.Println(binaryGap(22))
}
