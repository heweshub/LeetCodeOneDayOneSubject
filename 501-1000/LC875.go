package main

import (
	"sort"
)

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, v := range piles {
		if max < v {
			max = v
		}
	}
	return 1 + sort.Search(max-1, func(speed int) bool {
		speed++
		time := 0
		for _, v := range piles {
			time += (v + speed - 1) / speed
		}
		return time <= h
	})
}

// func main() {
// 	fmt.Println(minEatingSpeed([]int{312884470}, 312884469))
// }
