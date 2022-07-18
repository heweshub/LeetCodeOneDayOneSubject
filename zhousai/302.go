package main

import (
	"sort"
	"strconv"
)

func smallestTrimmedNumbers(nums []string, q [][]int) []int {
	ans := make([]int, len(q))
	for i := 0; i < len(q); i++ {
		n := make([]int, len(nums))
		m := make(map[int]int, len(nums))
		for j, v := range nums {
			n[j], _ = strconv.Atoi(v[len(v)-q[i][1]:])
			m[n[j]] = j
		}
		sort.Ints(n)
		ans[i] = m[n[q[i][0]]]
	}
	return ans
}

// func main() {

// }
