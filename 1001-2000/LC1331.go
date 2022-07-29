package main

import "sort"

func arrayRankTransform(arr []int) []int {
	arr2 := make([]int, len(arr))
	copy(arr2, arr)
	sort.Ints(arr2)

	m := map[int]int{}
	for i, v := range arr2 {
		m[v] = i + 1
	}
	// fmt.Println(m)
	for i := range arr {
		arr[i] = m[arr[i]]
	}
	return arr
}
