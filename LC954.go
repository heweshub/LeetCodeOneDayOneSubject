package main

import (
	"sort"
)

func canReorderDoubled(arr []int) bool {
	n := len(arr)
	cnt := make(map[int]int, n)
	for _, x := range arr {
		cnt[x]++
	}
	if cnt[0]%2 != 0 {
		return false
	}
	vals := make([]int, 0, len(cnt))
	for x := range cnt {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool { return abs(vals[i]) < abs(vals[j]) })
	for _, x := range vals {
		if cnt[2*x] < cnt[x] {
			return false
		}
		cnt[2*x] -= cnt[x]
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func main() {
// 	fmt.Println(canReorderDoubled([]int{2, 4, 0, 0, 8, 1}))
// }
