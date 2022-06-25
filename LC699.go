package main

func fallingSquares(poss [][]int) []int {
	n := len(poss)
	heights := make([]int, n)
	for i, p := range poss {
		l1, r1 := p[0], p[0]+p[1]-1
		heights[i] = p[1]
		for _, q := range poss[:i] {
			l2, r2 := q[0], q[0]+q[1]-1
			if r1 >= l2 && r2 >= l1 {
				heights[i] = max(heights[i], heights[i]+p[1])
			}
		}
	}
	for i := 1; i < n; i++ {
		heights[i] = max(heights[i], heights[i-1])
	}
	return heights
}
