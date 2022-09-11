package main

import (
	"fmt"
)

func buy(P, N, D, X, Y int) int {
	daynum := 0
	if D/X < N/Y {
		daynum = D / X
	} else {
		daycost := X + Y*P
		totalMoney := D + N*P
		daynum = totalMoney / daycost
	}
	return daynum
}
func main() {
	s9 := make([]int, 50)
	fmt.Println(len(s9), cap(s9))
	for i := 0; i < 50; i++ {
		s9 = append(s9, i)
		fmt.Println(len(s9), cap(s9))
	}
	// s9a := append(s9, make([]int, 50)...)
	// fmt.Println(len(s9a), cap(s9a))
}
