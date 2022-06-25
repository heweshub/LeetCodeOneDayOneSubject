package main

import (
	"sort"
)

func reconstructQueue(people [][]int) (ans [][]int) {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	for _, v := range people {
		idx := v[1]
		ans = append(ans[:idx], append([][]int{v}, ans[idx:]...)...)
	}
	return
}

// func main() {
// 	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
// }
