package main

import "math/rand"

type Solution struct {
	r, x, y float64
}

func Constructor(r float64, x float64, y float64) Solution {
	return Solution{r, x, y}
}

func (s *Solution) RandPoint() []float64 {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		if x*x+y*y < 1 {
			return []float64{s.x + x*s.r, s.y + y*s.r}
		}
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
