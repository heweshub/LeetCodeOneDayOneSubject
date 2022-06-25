package main

import "sort"

func maxConsecutive(bottom int, top int, special []int) (ans int) {
	special = append(special, []int{bottom - 1, top + 1}...)
	sort.Ints(special)
	for i := 1; i < len(special); i++ {
		ans = maxMax(ans, special[i]-special[i-1]-1)
	}
	return
}
func maxMax(x, y int) int {
	if x < y {
		return y
	}
	return x
}
