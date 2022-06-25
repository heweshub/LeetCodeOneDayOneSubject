package main

import (
	"sort"
)

func findKthNumber(m, n, k int) int {
	return sort.Search(m*n, func(x int) bool {
		count := x / n * n
		for i := x/n + 1; i <= m; i++ {
			count += x / i
		}
		return count >= k
	})
}

// func main() {
// 	fmt.Println(findKthNumber(2, 3, 6))
// }
