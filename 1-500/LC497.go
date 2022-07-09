package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	rects [][]int
	sum   []int
}

// func Constructor(rects [][]int) Solution {
// 	sum := make([]int, len(rects)+1)
// 	for i, r := range rects {
// 		a, b, x, y := r[0], r[1], r[2], r[3]
// 		sum[i+1] = sum[i] + (x-a+1)*(y-b+1)
// 	}
// 	return Solution{rects, sum}
// }

func (s *Solution) Pick() []int {
	k := rand.Intn(s.sum[len(s.sum)-1])
	rectIndex := sort.SearchInts(s.sum, k+1) - 1
	r := s.rects[rectIndex]
	a, b, y := r[0], r[1], r[3]
	da := (k - s.sum[rectIndex]) / (y - b + 1)
	db := (k - s.sum[rectIndex]) % (y - b + 1)
	return []int{a + da, b + db}
}

// func main() {
// 	s := Constructor([][]int{{-2, -2, -1, -1}, {1, 0, 3, 0}})
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(s.Pick())
// 	}
// }

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */