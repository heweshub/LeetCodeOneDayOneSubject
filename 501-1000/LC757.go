package main

import "sort"

func intersectionSizeTwo(inter [][]int) (ans int) {
	sort.Slice(inter, func(i, j int) bool {
		a, b := inter[i], inter[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})
	a, b := -1, -1
	for _, i := range inter {
		if i[0] > b {
			a = i[1] - 1
			b = i[1]
			ans += 2
		} else if i[0] > a {
			a = b
			b = i[1]
			ans++
		}
	}
	//
	// n, m := len(inter), 2
	// vals := make([][]int, n)
	// for i := n - 1; i >= 0; i-- {
	// 	for j, k := inter[i][0], len(vals[i]); k < m; k++ {
	// 		ans++
	// 		for p := i - 1; p >= 0 && inter[p][1] >= j; p-- {
	// 			vals[p] = append(vals[p], j)
	// 		}
	// 		j++
	// 	}
	// }
	return
}
