package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	// seed the random seed by time.Now()
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := randomPartition(a, l, r)
	// if result of func randomPartition equals index,
	// func of quickSelect directly return.
	// else binary search to solve this problem.
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	// randomly to select pivot for quicksort
	i := rand.Int()%(r-l+1) + l
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

// quicksort algorithm
// partition use the pivot to apart the array a.
// return the index of pivot
func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

// func main() {
// 	findKthLargest([]int{3, 2, 1, 5, 6, 4}, 4)
// }
