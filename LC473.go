package main

import (
	"sort"
)

func makesquare(match []int) bool {
	total := 0
	for _, l := range match {
		total += l
	}
	if total%4 != 0 {
		return false
	}
	sort.Sort(sort.Reverse(sort.IntSlice(match)))
	edges := [4]int{}
	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(match) {
			return true
		}
		for i := range edges {
			edges[i] += match[idx]
			if edges[i] <= total/4 && dfs(idx+1) {
				return true
			}
			edges[i] -= match[idx]
		}
		return false
	}
	return dfs(0)
}

// func main() {
// 	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
// }
