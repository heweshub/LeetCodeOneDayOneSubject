package main

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	cnt := make([]int, n)
	for _, r := range roads {
		cnt[r[0]]++
		cnt[r[1]]++
	}
	sort.Ints(cnt)
	var res int
	for i, c := range cnt {
		res += (i + 1) * c
	}
	return int64(res)
}
