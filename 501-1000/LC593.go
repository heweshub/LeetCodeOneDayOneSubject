package main

import (
	"sort"
)

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	a1 := calcul_length(p1, p2)
	a2 := calcul_length(p1, p3)
	a3 := calcul_length(p1, p4)
	b1 := calcul_length(p2, p3)
	b2 := calcul_length(p2, p4)
	c1 := calcul_length(p3, p4)
	res := []int{a1, a2, a3, b1, b2, c1}
	sort.Ints(res)
	var size int
	for i := range res {
		if i == 0 {
			size = res[i]
		} else if i > 0 && i < 4 {
			if size != res[i] {
				return false
			}
		} else {
			if size+size == res[i] && size != 0 {
				return true
			}
		}
	}
	return false
}

func calcul_length(n1, n2 []int) int {
	x1, y1 := n1[0], n1[1]
	x2, y2 := n2[0], n2[1]
	return (y1-y2)*(y1-y2) + (x1-x2)*(x1-x2)
}
